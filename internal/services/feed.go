package services

import (
	"github.com/MehulxBuilds/RSSFeed/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

type FeedService struct {
	config config.Config
	db *pgxpool.Pool
}

func NewFeedService(cfg config.Config, db *pgxpool.Pool) *FeedService {
	return &FeedService{
		config: cfg,
		db: db,
	}
}