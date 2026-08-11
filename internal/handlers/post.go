package handlers

import (
	"github.com/MehulxBuilds/RSSFeed/internal/config"
	"github.com/MehulxBuilds/RSSFeed/internal/middleware"
	"github.com/MehulxBuilds/RSSFeed/internal/services"
	"github.com/jackc/pgx/v5/pgxpool"
	"net/http"
	"strconv"
)

type PostHandler struct {
	config      config.Config
	authService *services.PostService
	db          *pgxpool.Pool
}

func (h *PostHandler) Get(w http.ResponseWriter, r *http.Request) {
	u, _ := middleware.UserFromContext(r.Context())
	limit := 10
	if raw := r.URL.Query().Get("limit"); raw != "" {
		n, err := strconv.Atoi(raw)
		if err != nil || n < 1 || n > 1000 {
			respondError(w, 400, "limit must be between 1 and 1000")
			return
		}
		limit = n
	}
	posts, err := h.authService.ForUser(r.Context(), u.ID, limit)
	if err != nil {
		respondError(w, 500, "Couldn't get posts for user")
		return
	}
	respondJSON(w, 200, posts)
}

func NewPostHandler(cfg config.Config, db *pgxpool.Pool, authService *services.PostService) *PostHandler {
	return &PostHandler{
		config:      cfg,
		db:          db,
		authService: authService,
	}
}
