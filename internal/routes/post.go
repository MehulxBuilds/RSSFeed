package routes

import (
	"github.com/MehulxBuilds/RSSFeed/internal/handlers"
	"github.com/MehulxBuilds/RSSFeed/internal/middleware"
	"github.com/go-chi/chi"
)

type PostRouteDependencies struct {
	PostHandler    *handlers.PostHandler
	PostMiddleware *middleware.AuthMiddleware
}

func RegisterPostRoutes(router *chi.Mux, deps PostRouteDependencies) {
	router.Route("/posts", func(r chi.Router) {
		r.Use(deps.PostMiddleware.Authenticate)
		r.Get("/", deps.PostHandler.Get)
	})
}
