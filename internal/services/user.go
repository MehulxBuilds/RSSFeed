package services

import (
	"github.com/MehulxBuilds/RSSFeed/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthService struct {
	config config.Config
	db *pgxpool.Pool
}

func NewUserService(cfg config.Config, db *pgxpool.Pool) *AuthService {
	return &AuthService{
		config: cfg,
		db: db,
	}
}