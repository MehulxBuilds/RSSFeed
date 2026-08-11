package handlers

import (
	"encoding/json"
	"github.com/MehulxBuilds/RSSFeed/internal/config"
	"github.com/MehulxBuilds/RSSFeed/internal/middleware"
	"github.com/MehulxBuilds/RSSFeed/internal/services"
	"github.com/go-chi/chi"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"net/http"
)

type FeedFollowsHandler struct {
	config      config.Config
	authService *services.FeedFollowService
	db          *pgxpool.Pool
}

func NewFeedFollowHandler(cfg config.Config, db *pgxpool.Pool, authService *services.FeedFollowService) *FeedFollowsHandler {
	return &FeedFollowsHandler{
		config:      cfg,
		db:          db,
		authService: authService,
	}
}

func (h *FeedFollowsHandler) Create(w http.ResponseWriter, r *http.Request) {
	u, ok := middleware.UserFromContext(r.Context())
	if !ok {
		respondError(w, 401, "Unauthorized")
		return
	}
	var body struct {
		FeedID string `json:"feed_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body.FeedID == "" {
		respondError(w, 400, "feed_id is required")
		return
	}
	var feedID pgtype.UUID
	if err := feedID.Scan(body.FeedID); err != nil || !feedID.Valid {
		respondError(w, http.StatusBadRequest, "Invalid feed ID")
		return
	}
	f, err := h.authService.Create(r.Context(), u.ID, body.FeedID)
	if err != nil {
		respondError(w, 500, "Couldn't create feed follow")
		return
	}
	respondJSON(w, http.StatusOK, f)
}
func (h *FeedFollowsHandler) Get(w http.ResponseWriter, r *http.Request) {
	u, _ := middleware.UserFromContext(r.Context())
	f, err := h.authService.ForUser(r.Context(), u.ID)
	if err != nil {
		respondError(w, 500, "Couldn't get feed follows")
		return
	}
	respondJSON(w, 200, f)
}
func (h *FeedFollowsHandler) Delete(w http.ResponseWriter, r *http.Request) {
	u, _ := middleware.UserFromContext(r.Context())
	id := chi.URLParam(r, "feedFollowID")
	var feedFollowID pgtype.UUID
	if err := feedFollowID.Scan(id); err != nil || !feedFollowID.Valid {
		respondError(w, http.StatusBadRequest, "Invalid feed follow ID")
		return
	}
	if err := h.authService.Delete(r.Context(), u.ID, id); err != nil {
		respondError(w, 404, "Feed follow not found")
		return
	}
	respondJSON(w, 200, struct{}{})
}
