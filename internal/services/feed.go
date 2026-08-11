package services

import (
	"context"
	"time"

	"github.com/MehulxBuilds/RSSFeed/internal/config"
	"github.com/MehulxBuilds/RSSFeed/internal/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type FeedService struct {
	config config.Config
	db     *pgxpool.Pool
}

func NewFeedService(cfg config.Config, db *pgxpool.Pool) *FeedService {
	return &FeedService{
		config: cfg,
		db:     db,
	}
}

func (s *FeedService) Create(ctx context.Context, userID, name, url string) (models.Feed, models.FeedFollow, error) {
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return models.Feed{}, models.FeedFollow{}, err
	}
	defer tx.Rollback(ctx)
	now := time.Now().UTC()
	var feed models.Feed
	err = tx.QueryRow(ctx, `INSERT INTO feeds (id,created_at,updated_at,name,url,user_id) VALUES ($1,$2,$3,$4,$5,$6) RETURNING id,created_at,updated_at,name,url,user_id,last_fetched_at`, uuid.New(), now, now, name, url, userID).
		Scan(&feed.ID, &feed.CreatedAt, &feed.UpdatedAt, &feed.Name, &feed.URL, &feed.UserID, &feed.LastFetchedAt)
	if err != nil {
		return models.Feed{}, models.FeedFollow{}, err
	}
	var follow models.FeedFollow
	err = tx.QueryRow(ctx, `INSERT INTO feed_follows (id,created_at,updated_at,user_id,feed_id) VALUES ($1,$2,$3,$4,$5) RETURNING id,created_at,updated_at,user_id,feed_id`, uuid.New(), now, now, userID, feed.ID).
		Scan(&follow.ID, &follow.CreatedAt, &follow.UpdatedAt, &follow.UserID, &follow.FeedID)
	if err != nil {
		return models.Feed{}, models.FeedFollow{}, err
	}
	if err = tx.Commit(ctx); err != nil {
		return models.Feed{}, models.FeedFollow{}, err
	}
	return feed, follow, nil
}

func (s *FeedService) All(ctx context.Context) ([]models.Feed, error) {
	rows, err := s.db.Query(ctx, `SELECT id,created_at,updated_at,name,url,user_id,last_fetched_at FROM feeds ORDER BY created_at DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	feeds := make([]models.Feed, 0)
	for rows.Next() {
		var f models.Feed
		if err := rows.Scan(&f.ID, &f.CreatedAt, &f.UpdatedAt, &f.Name, &f.URL, &f.UserID, &f.LastFetchedAt); err != nil {
			return nil, err
		}
		feeds = append(feeds, f)
	}
	return feeds, rows.Err()
}
