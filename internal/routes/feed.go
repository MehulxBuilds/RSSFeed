package routes

import (
	"github.com/MehulxBuilds/RSSFeed/internal/handlers"
	"github.com/MehulxBuilds/RSSFeed/internal/middleware"
	"github.com/go-chi/chi"
)

type FeedRouteDependencies struct {
	FeedHandler    *handlers.FeedsHandler
	FeedMiddleware *middleware.AuthMiddleware
}

func RegisterFeedRoutes(router *chi.Mux, deps FeedRouteDependencies) {
	router.Route("/feeds", func(r chi.Router) {
		r.Get("/", deps.FeedHandler.GetAll)
		r.With(deps.FeedMiddleware.Authenticate).Post("/", deps.FeedHandler.Create)
	})
}
