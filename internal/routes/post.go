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
	router.Route("/post", func(r chi.Router) {
		
	})
}
