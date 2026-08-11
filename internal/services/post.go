package services

import (
	"github.com/MehulxBuilds/RSSFeed/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostService struct {
	config config.Config
	db *pgxpool.Pool
}

func NewPostService(cfg config.Config, db *pgxpool.Pool) *PostService {
	return &PostService{
		config: cfg,
		db: db,
	}
}