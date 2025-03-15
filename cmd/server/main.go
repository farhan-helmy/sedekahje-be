package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/farhan-helmy/sedekahje-be/internal/config"
	"github.com/farhan-helmy/sedekahje-be/internal/db"
	"github.com/farhan-helmy/sedekahje-be/internal/routes"
	"github.com/farhan-helmy/sedekahje-be/internal/utils"
)

func main() {

	cfg := config.LoadConfig()

	mongoClient := db.ConnectDB(cfg.MongoURI)

	router := mux.NewRouter()

	routes.SetupRoutes(router, mongoClient)

	router.Use(utils.LoggingMiddleware)

	// Start server
	log.Println("Server running on :3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
