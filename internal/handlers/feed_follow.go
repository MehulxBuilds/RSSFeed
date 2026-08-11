package handlers

import (
	"github.com/MehulxBuilds/RSSFeed/internal/config"
	"github.com/MehulxBuilds/RSSFeed/internal/services"
	"github.com/jackc/pgx/v5/pgxpool"
)

type FeedFollowsHandler struct {
	config config.Config
	authService *services.FeedFollowService
	db *pgxpool.Pool
}

func NewFeedFollowHandler(cfg config.Config, db *pgxpool.Pool, authService *services.FeedFollowService) *FeedFollowsHandler {
	return &FeedFollowsHandler{
		config: cfg,
		db: db,
		authService: authService,
	}
}