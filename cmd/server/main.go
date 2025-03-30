package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"

	"github.com/common-nighthawk/go-figure"
	"github.com/farhan-helmy/sedekahje-be/internal/config"
	"github.com/farhan-helmy/sedekahje-be/internal/db"
	"github.com/farhan-helmy/sedekahje-be/internal/routes"
	"github.com/farhan-helmy/sedekahje-be/internal/utils"
	"github.com/throttled/throttled/v2"
	"github.com/throttled/throttled/v2/store/memstore"
)

func gracefulShutdown(apiServer *http.Server, done chan bool) {
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Listen for the interrupt signal.
	<-ctx.Done()

	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := apiServer.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown with error: %v", err)
	}

	log.Println("Server exiting")

	// Notify the main goroutine that the shutdown is complete
	done <- true
}

func main() {
	cfg := config.LoadConfig()

	mongoClient := db.ConnectDB(cfg.MongoURI)

	router := mux.NewRouter()

	// Create a done channel to signal when the shutdown is complete
	done := make(chan bool, 1)

	apiServer := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go gracefulShutdown(apiServer, done)

	routes.SetupRoutes(router, mongoClient)

	router.Use(utils.LoggingMiddleware)

	store, err := memstore.NewCtx(65536)
	if err != nil {
		log.Fatal(err)
	}

	quota := throttled.RateQuota{
		MaxRate:  throttled.PerMin(20),
		MaxBurst: 5,
	}

	rateLimiter, err := throttled.NewGCRARateLimiterCtx(store, quota)
	if err != nil {
		log.Fatal(err)
	}

	httpRateLimiter := throttled.HTTPRateLimiterCtx{
		RateLimiter: rateLimiter,
		VaryBy:      &throttled.VaryBy{Path: true},
	}

	// Apply rate limiting middleware to the router
	router.Use(httpRateLimiter.RateLimit)

	// Start the server in a goroutine so that it doesn't block
	go func() {
		myFigure := figure.NewFigure("Sedekahje", "", true)
		myFigure.Print()
		log.Println("Server running on :8080")
		if err := apiServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe(): %s", err)
		}
	}()

	// Wait for the shutdown signal
	<-done
	log.Println("Server has shut down gracefully")

	// Close the MongoDB connection
	if err := mongoClient.Disconnect(context.Background()); err != nil {
		log.Printf("Error disconnecting from MongoDB: %v", err)
	} else {
		log.Println("Disconnected from MongoDB")
	}
}
