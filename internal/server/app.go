package server

import (
	"github.com/MehulxBuilds/RSSFeed/internal/config"
	"github.com/go-chi/chi"
	"github.com/jackc/pgx/v5/pgxpool"
)

func New(cfg config.Config, db *pgxpool.Pool, router *chi.Mux) error {
	return nil
}