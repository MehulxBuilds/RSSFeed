package services

import (
	"context"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type rssFeed struct {
	Channel struct {
		Items []rssItem `xml:"item"`
	} `xml:"channel"`
}
type rssItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func StartScraping(ctx context.Context, db *pgxpool.Pool, concurrency int, interval time.Duration) {
	if concurrency < 1 {
		concurrency = 1
	}
	run := func() { scrapeBatch(ctx, db, concurrency) }
	run()
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			run()
		}
	}
}

func scrapeBatch(ctx context.Context, db *pgxpool.Pool, limit int) {
	rows, err := db.Query(ctx, `SELECT id,name,url FROM feeds ORDER BY last_fetched_at ASC NULLS FIRST LIMIT $1`, limit)
	if err != nil {
		log.Printf("get feeds to fetch: %v", err)
		return
	}
	type feed struct{ id, name, url string }
	feeds := make([]feed, 0, limit)
	for rows.Next() {
		var f feed
		if err := rows.Scan(&f.id, &f.name, &f.url); err != nil {
			log.Printf("scan feed: %v", err)
			continue
		}
		feeds = append(feeds, f)
	}
	rows.Close()
	var wg sync.WaitGroup
	for _, f := range feeds {
		wg.Add(1)
		go func() { defer wg.Done(); scrapeOne(ctx, db, f.id, f.name, f.url) }()
	}
	wg.Wait()
}

func scrapeOne(ctx context.Context, db *pgxpool.Pool, id, name, url string) {
	if _, err := db.Exec(ctx, `UPDATE feeds SET last_fetched_at=NOW(),updated_at=NOW() WHERE id=$1`, id); err != nil {
		log.Printf("mark feed %s fetched: %v", name, err)
		return
	}
	feed, err := fetchRSS(ctx, url)
	if err != nil {
		log.Printf("fetch feed %s: %v", name, err)
		return
	}
	for _, item := range feed.Channel.Items {
		published := parsePublishedAt(item.PubDate)
		_, err = db.Exec(ctx, `INSERT INTO posts (id,created_at,updated_at,title,url,description,published_at,feed_id) VALUES ($1,NOW(),NOW(),$2,$3,$4,$5,$6)`, uuid.New(), item.Title, item.Link, nullableString(item.Description), published, id)
		if err != nil && !isUniqueViolation(err) {
			log.Printf("create post for %s: %v", name, err)
		}
	}
}

func fetchRSS(ctx context.Context, url string) (*rssFeed, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("User-Agent", "RSSFeed/1.0")
	client := http.Client{Timeout: 10 * time.Second}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected HTTP status %s", response.Status)
	}
	reader := io.LimitReader(response.Body, 10<<20)
	var feed rssFeed
	if err := xml.NewDecoder(reader).Decode(&feed); err != nil {
		return nil, err
	}
	return &feed, nil
}

func parsePublishedAt(value string) *time.Time {
	for _, layout := range []string{time.RFC1123Z, time.RFC1123, time.RFC822Z, time.RFC822, time.RFC3339, time.RFC3339Nano, "Mon, 2 Jan 2006 15:04:05 -0700"} {
		if parsed, err := time.Parse(layout, value); err == nil {
			parsed = parsed.UTC()
			return &parsed
		}
	}
	return nil
}
func nullableString(value string) *string {
	if value == "" {
		return nil
	}
	return &value
}
func isUniqueViolation(err error) bool {
	var pgErr *pgconn.PgError
	return errors.As(err, &pgErr) && pgErr.Code == "23505"
}
