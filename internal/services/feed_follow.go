package services

import (
	"context"
	"fmt"
	"time"

	"github.com/MehulxBuilds/RSSFeed/internal/config"
	"github.com/MehulxBuilds/RSSFeed/internal/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type FeedFollowService struct {
	config config.Config
	db     *pgxpool.Pool
}

func NewFeedFollowService(cfg config.Config, db *pgxpool.Pool) *FeedFollowService {
	return &FeedFollowService{
		config: cfg,
		db:     db,
	}
}

func (s *FeedFollowService) Create(ctx context.Context, userID, feedID string) (models.FeedFollow, error) {
	now := time.Now().UTC()
	var f models.FeedFollow
	err := s.db.QueryRow(ctx, `INSERT INTO feed_follows (id,created_at,updated_at,user_id,feed_id) VALUES ($1,$2,$3,$4,$5) RETURNING id,created_at,updated_at,user_id,feed_id`, uuid.New(), now, now, userID, feedID).Scan(&f.ID, &f.CreatedAt, &f.UpdatedAt, &f.UserID, &f.FeedID)
	return f, err
}
func (s *FeedFollowService) ForUser(ctx context.Context, userID string) ([]models.FeedFollow, error) {
	rows, err := s.db.Query(ctx, `SELECT id,created_at,updated_at,user_id,feed_id FROM feed_follows WHERE user_id=$1 ORDER BY created_at DESC`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	result := make([]models.FeedFollow, 0)
	for rows.Next() {
		var f models.FeedFollow
		if err := rows.Scan(&f.ID, &f.CreatedAt, &f.UpdatedAt, &f.UserID, &f.FeedID); err != nil {
			return nil, err
		}
		result = append(result, f)
	}
	return result, rows.Err()
}
func (s *FeedFollowService) Delete(ctx context.Context, userID, id string) error {
	tag, err := s.db.Exec(ctx, `DELETE FROM feed_follows WHERE id=$1 AND user_id=$2`, id, userID)
	if err == nil && tag.RowsAffected() == 0 {
		return fmt.Errorf("feed follow not found")
	}
	return err
}
