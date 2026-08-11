package middleware

import (
	"context"
	"github.com/MehulxBuilds/RSSFeed/internal/config"
	"github.com/MehulxBuilds/RSSFeed/internal/models"
	"github.com/MehulxBuilds/RSSFeed/internal/services"
	"net/http"
	"strings"
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

type contextKey string

const userKey contextKey = "authenticated-user"

func UserFromContext(ctx context.Context) (models.User, bool) {
	u, ok := ctx.Value(userKey).(models.User)
	return u, ok
}

func (m *AuthMiddleware) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		parts := strings.Fields(r.Header.Get("Authorization"))
		if len(parts) != 2 || parts[0] != "ApiKey" || parts[1] == "" {
			writeAuthError(w, http.StatusUnauthorized, "Couldn't find api key")
			return
		}
		user, err := m.authService.ByAPIKey(r.Context(), parts[1])
		if err != nil {
			writeAuthError(w, http.StatusNotFound, "Couldn't get user")
			return
		}
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), userKey, user)))
	})
}

func writeAuthError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, _ = w.Write([]byte(`{"error":"` + message + `"}`))
}
