package handlers

import (
	"encoding/json"
	"github.com/MehulxBuilds/RSSFeed/internal/config"
	"github.com/MehulxBuilds/RSSFeed/internal/middleware"
	"github.com/MehulxBuilds/RSSFeed/internal/services"
	"github.com/jackc/pgx/v5/pgxpool"
	"net/http"
)

type UserHandler struct {
	config      config.Config
	authService *services.AuthService
	db          *pgxpool.Pool
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body.Name == "" {
		respondError(w, http.StatusBadRequest, "name is required")
		return
	}
	u, err := h.authService.Create(r.Context(), body.Name)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Couldn't create user")
		return
	}
	respondJSON(w, http.StatusOK, u)
}
func (h *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	u, ok := middleware.UserFromContext(r.Context())
	if !ok {
		respondError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}
	respondJSON(w, http.StatusOK, u)
}

func NewUserHandler(cfg config.Config, db *pgxpool.Pool, authService *services.AuthService) *UserHandler {
	return &UserHandler{
		config:      cfg,
		db:          db,
		authService: authService,
	}
}
