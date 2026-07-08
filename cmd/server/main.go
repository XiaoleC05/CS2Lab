package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/XiaoleC05/CS2Lab/internal/config"
	"github.com/XiaoleC05/CS2Lab/internal/db"
	"github.com/XiaoleC05/CS2Lab/internal/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize database
	if err := db.Init(cfg); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	// Run migrations
	log.Println("Running database migrations...")
	if err := db.RunMigrations("001_init"); err != nil {
		log.Fatalf("Failed to run init migration: %v", err)
	}

	if err := db.RunMigrations("002_seed_data"); err != nil {
		log.Printf("Warning: seed data migration failed: %v (this is OK if data already exists)", err)
	}

	// Set up Gin router
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// Configure CORS
	corsConfig := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-User-ID"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	r.Use(cors.New(corsConfig))

	// Serve static files (images)
	r.Static("/images", cfg.CS2LabStaticDir)

	// Initialize handlers
	mapHandler := handler.NewMapHandler()
	lineupHandler := handler.NewLineupHandler()
	favoriteHandler := handler.NewFavoriteHandler()
	noteHandler := handler.NewNoteHandler()

	// Public routes (no authentication required)
	api := r.Group("/api")
	{
		api.GET("/health", handler.HealthHandler)
		api.GET("/maps", mapHandler.GetAll)
		api.GET("/maps/:id", mapHandler.GetByID)
		api.GET("/lineups", lineupHandler.GetFiltered)
		api.GET("/lineups/:id", lineupHandler.GetByID)
	}

	// Protected routes (authentication required)
	protected := r.Group("/api")
	protected.Use(handler.AuthMiddleware(cfg))
	{
		protected.GET("/favorites", favoriteHandler.GetByUser)
		protected.POST("/favorites", favoriteHandler.Add)
		protected.DELETE("/favorites/:lineupId", favoriteHandler.Remove)
		protected.GET("/notes/:lineupId", noteHandler.GetByLineup)
		protected.PUT("/notes/:lineupId", noteHandler.Upsert)
	}

	// Set up graceful shutdown
	srv := &http.Server{
		Addr:    ":" + cfg.CS2LabPort,
		Handler: r,
	}

	go func() {
		log.Printf("CS2Lab server starting on port %s", cfg.CS2LabPort)
		log.Printf("Gateway mode: %v", cfg.OxeliaGatewayMode)
		log.Printf("Static files directory: %s", cfg.CS2LabStaticDir)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	// Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited gracefully")
}
