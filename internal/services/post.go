package services

import (
	"context"
	"github.com/MehulxBuilds/RSSFeed/internal/config"
	"github.com/MehulxBuilds/RSSFeed/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostService struct {
	config config.Config
	db     *pgxpool.Pool
}

func NewPostService(cfg config.Config, db *pgxpool.Pool) *PostService {
	return &PostService{
		config: cfg,
		db:     db,
	}
}

func (s *PostService) ForUser(ctx context.Context, userID string, limit int) ([]models.Post, error) {
	rows, err := s.db.Query(ctx, `SELECT p.id,p.created_at,p.updated_at,p.title,p.url,p.description,p.published_at,p.feed_id
		FROM posts p JOIN feed_follows ff ON ff.feed_id=p.feed_id
		WHERE ff.user_id=$1 ORDER BY p.published_at DESC NULLS LAST, p.created_at DESC LIMIT $2`, userID, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	posts := make([]models.Post, 0)
	for rows.Next() {
		var p models.Post
		if err := rows.Scan(&p.ID, &p.CreatedAt, &p.UpdatedAt, &p.Title, &p.URL, &p.Description, &p.PublishedAt, &p.FeedID); err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}
	return posts, rows.Err()
}
