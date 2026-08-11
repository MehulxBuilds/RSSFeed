package handlers

import (
	"github.com/MehulxBuilds/RSSFeed/internal/config"
	"github.com/MehulxBuilds/RSSFeed/internal/services"
	"github.com/jackc/pgx/v5/pgxpool"
)

type FeedsHandler struct {
	config config.Config
	authService *services.FeedService
	db *pgxpool.Pool
}

func NewFeedHandler(cfg config.Config, db *pgxpool.Pool, authService *services.FeedService) *FeedsHandler {
	return &FeedsHandler{
		config: cfg,
		db: db,
		authService: authService,
	}
}