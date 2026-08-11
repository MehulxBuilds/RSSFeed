package services

import (
	"context"
	"github.com/MehulxBuilds/RSSFeed/internal/config"
	"github.com/MehulxBuilds/RSSFeed/internal/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type AuthService struct {
	config config.Config
	db     *pgxpool.Pool
}

func NewUserService(cfg config.Config, db *pgxpool.Pool) *AuthService {
	return &AuthService{
		config: cfg,
		db:     db,
	}
}

func (s *AuthService) Create(ctx context.Context, name string) (models.User, error) {
	now := time.Now().UTC()
	var user models.User
	err := s.db.QueryRow(ctx, `INSERT INTO users (id, created_at, updated_at, name) VALUES ($1,$2,$3,$4) RETURNING id, created_at, updated_at, name, api_key`, uuid.New(), now, now, name).
		Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt, &user.Name, &user.APIKey)
	return user, err
}

func (s *AuthService) ByAPIKey(ctx context.Context, apiKey string) (models.User, error) {
	var user models.User
	err := s.db.QueryRow(ctx, `SELECT id, created_at, updated_at, name, api_key FROM users WHERE api_key=$1`, apiKey).
		Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt, &user.Name, &user.APIKey)
	return user, err
}
