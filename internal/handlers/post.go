package handlers

import (
	"github.com/MehulxBuilds/RSSFeed/internal/config"
	"github.com/MehulxBuilds/RSSFeed/internal/services"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostHandler struct {
	config config.Config
	authService *services.PostService
	db *pgxpool.Pool
}

func NewPostHandler(cfg config.Config, db *pgxpool.Pool, authService *services.PostService) *PostHandler {
	return &PostHandler{
		config: cfg,
		db: db,
		authService: authService,
	}
}
