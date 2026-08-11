package main

import (
	"context"
	"log"
	"net/http"

	"github.com/MehulxBuilds/RSSFeed/internal/config"
	"github.com/MehulxBuilds/RSSFeed/internal/database"
	"github.com/MehulxBuilds/RSSFeed/internal/server"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func main() {

	// Load Context
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Load config: %v", err)
	}

	// Base Context
	ctx := context.Background()

	db, err := database.NewPool(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Connect Database: %v", err)
	}

	defer db.Close()
	
	def_router := chi.NewRouter()

	def_router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// Root Route
	def_router.Get("/", server.RootRoute)

	// Heath Check
	def_router.Get("/health", server.HealthCheck)

	v1Router := chi.NewRouter()
	def_router.Mount("/v1", v1Router)

	// Make Chi App ( This will handle everything under the hood )
	err = server.New(cfg, db, v1Router)
	if err != nil {
		log.Fatalf("Server Registration Error: %v", err)
	}

	srv := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: def_router,
	}

	log.Printf("Serving on port: %s\n", cfg.Port)
	log.Fatal(srv.ListenAndServe())
}
