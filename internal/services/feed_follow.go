package services

import (
	"github.com/MehulxBuilds/RSSFeed/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

type FeedFollowService struct {
	config config.Config
	db *pgxpool.Pool
}

func NewFeedFollowService(cfg config.Config, db *pgxpool.Pool) *FeedFollowService {
	return &FeedFollowService{
		config: cfg,
		db: db,
	}
}