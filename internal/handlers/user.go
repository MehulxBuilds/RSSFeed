package handlers

import (
	"github.com/MehulxBuilds/RSSFeed/internal/config"
	"github.com/MehulxBuilds/RSSFeed/internal/services"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserHandler struct {
	config config.Config
	authService *services.AuthService
	db *pgxpool.Pool
}

func NewUserHandler(cfg config.Config, db *pgxpool.Pool, authService *services.AuthService) *UserHandler {
	return &UserHandler{
		config: cfg,
		db: db,
		authService: authService,
	}
}
