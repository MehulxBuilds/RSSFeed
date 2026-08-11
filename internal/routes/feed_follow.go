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
	router.Route("/feed-follow", func(r chi.Router) {
		
	})
}
