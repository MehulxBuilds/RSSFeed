package middleware

import (
	"github.com/MehulxBuilds/RSSFeed/internal/config"
	"github.com/MehulxBuilds/RSSFeed/internal/services"
)

type AuthMiddleware struct {
	config config.Config

	authService *services.AuthService
}

func NewAuthMiddleware(
	cfg config.Config,
	authService *services.AuthService,
) *AuthMiddleware {
	return &AuthMiddleware{
		config:      cfg,
		authService: authService,
	}
}