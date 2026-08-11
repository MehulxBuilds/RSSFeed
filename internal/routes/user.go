package routes

import (
	"github.com/MehulxBuilds/RSSFeed/internal/handlers"
	"github.com/MehulxBuilds/RSSFeed/internal/middleware"
	"github.com/go-chi/chi"
)

type AuthRouteDependencies struct {
	AuthHandler    *handlers.UserHandler
	AuthMiddleware *middleware.AuthMiddleware
}

func RegisterUserRoutes(router *chi.Mux, deps AuthRouteDependencies) {
	router.Route("/users", func(r chi.Router) {
		r.Post("/", deps.AuthHandler.Create)
		r.With(deps.AuthMiddleware.Authenticate).Get("/", deps.AuthHandler.Get)
	})
}
