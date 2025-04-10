package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"rockstar-corp/internal/config"
	"rockstar-corp/internal/infrastructure/db"
	"rockstar-corp/internal/interfaces/http/router"
)

func main() {
	// Load environment variables and configs
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// Initialize database connection (PostgreSQL example)
	dbConn, err := db.NewPostgres(cfg.Database)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer dbConn.Close()

	// Initialize HTTP router
	r := router.New(cfg, dbConn)

	srv := &http.Server{
		Addr:    ":" + cfg.HTTP.Port,
		Handler: r,
	}

	// Start server in a goroutine
	go func() {
		log.Printf("🚀 Server is running on port %s", cfg.HTTP.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("🛑 Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("✅ Server exited properly")
}
