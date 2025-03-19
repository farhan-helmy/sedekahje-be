package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/farhan-helmy/sedekahje-be/internal/config"
	"github.com/farhan-helmy/sedekahje-be/internal/db"
	"github.com/farhan-helmy/sedekahje-be/internal/routes"
	"github.com/farhan-helmy/sedekahje-be/internal/utils"
	"github.com/throttled/throttled/v2"
	"github.com/throttled/throttled/v2/store/memstore"
)

func main() {

	cfg := config.LoadConfig()

	mongoClient := db.ConnectDB(cfg.MongoURI)

	router := mux.NewRouter()

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

	// Start server
	log.Println("Server running on :3000")
	log.Fatal(http.ListenAndServe(":3000", httpRateLimiter.RateLimit(router)))
}
