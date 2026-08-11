package routes

import (
	"github.com/MehulxBuilds/RSSFeed/internal/handlers"
	"github.com/MehulxBuilds/RSSFeed/internal/middleware"
	"github.com/go-chi/chi"
)

type FeedFollowRouteDependencies struct {
	FeedFollowHandler    *handlers.FeedFollowsHandler
	FeedFollowMiddleware *middleware.AuthMiddleware
}

func RegisterFeedFollowRoutes(router *chi.Mux, deps FeedFollowRouteDependencies) {
	router.Route("/feed_follows", func(r chi.Router) {
		r.Use(deps.FeedFollowMiddleware.Authenticate)
		r.Get("/", deps.FeedFollowHandler.Get)
		r.Post("/", deps.FeedFollowHandler.Create)
		r.Delete("/{feedFollowID}", deps.FeedFollowHandler.Delete)
	})
}
