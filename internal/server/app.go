package server

import (
	"github.com/MehulxBuilds/RSSFeed/internal/config"
	"github.com/MehulxBuilds/RSSFeed/internal/handlers"
	"github.com/MehulxBuilds/RSSFeed/internal/middleware"
	"github.com/MehulxBuilds/RSSFeed/internal/routes"
	"github.com/MehulxBuilds/RSSFeed/internal/services"
	"github.com/go-chi/chi"
	"github.com/jackc/pgx/v5/pgxpool"
)

func New(cfg config.Config, db *pgxpool.Pool, router *chi.Mux) error {

	// Register Services here
	authService := services.NewUserService(cfg, db)
	postService := services.NewPostService(cfg, db)
	feedService := services.NewFeedService(cfg, db)
	feedfollowService := services.NewFeedFollowService(cfg, db)

	// Register Handlers here
	userHandler := handlers.NewUserHandler(cfg, db, authService)
	postHandler := handlers.NewPostHandler(cfg, db, postService)
	feedHandler := handlers.NewFeedHandler(cfg, db, feedService)
	feedfollowHandler := handlers.NewFeedFollowHandler(cfg, db, feedfollowService)

	// Register Middleware here
	authMiddleware := middleware.NewAuthMiddleware(cfg, authService)

	// register routes here
	routes.RegisterUserRoutes(router, routes.AuthRouteDependencies{
		AuthHandler: userHandler,
		AuthMiddleware: authMiddleware,
	})
	routes.RegisterPostRoutes(router, routes.PostRouteDependencies{
		PostHandler: postHandler,
		PostMiddleware: authMiddleware,
	})
	routes.RegisterFeedRoutes(router, routes.FeedRouteDependencies{
		FeedHandler: feedHandler,
		FeedMiddleware: authMiddleware,
	})
	routes.RegisterFeedFollowRoutes(router, routes.FeedFollowRouteDependencies{
		FeedFollowHandler: feedfollowHandler,
		FeedFollowMiddleware: authMiddleware,
	})

	return nil
}
