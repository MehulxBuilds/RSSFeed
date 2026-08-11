package handlers

import (
	"encoding/json"
	"github.com/MehulxBuilds/RSSFeed/internal/config"
	"github.com/MehulxBuilds/RSSFeed/internal/middleware"
	"github.com/MehulxBuilds/RSSFeed/internal/services"
	"github.com/jackc/pgx/v5/pgxpool"
	"net/http"
)

type FeedsHandler struct {
	config      config.Config
	feedService *services.FeedService
	db          *pgxpool.Pool
}

func NewFeedHandler(cfg config.Config, db *pgxpool.Pool, feedService *services.FeedService) *FeedsHandler {
	return &FeedsHandler{
		config:      cfg,
		db:          db,
		feedService: feedService,
	}
}

func (h *FeedsHandler) Create(w http.ResponseWriter, r *http.Request) {
	u, ok := middleware.UserFromContext(r.Context())
	if !ok {
		respondError(w, 401, "Unauthorized")
		return
	}
	var body struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body.Name == "" || body.URL == "" {
		respondError(w, 400, "name and url are required")
		return
	}
	f, ff, err := h.feedService.Create(r.Context(), u.ID, body.Name, body.URL)
	if err != nil {
		respondError(w, 500, "Couldn't create feed")
		return
	}
	respondJSON(w, http.StatusOK, map[string]any{"feed": f, "feed_follow": ff})
}
func (h *FeedsHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	feeds, err := h.feedService.All(r.Context())
	if err != nil {
		respondError(w, 500, "Couldn't get feeds")
		return
	}
	respondJSON(w, 200, feeds)
}
