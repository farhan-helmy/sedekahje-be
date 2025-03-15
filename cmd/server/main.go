package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/farhan-helmy/sedekahje-be/internal/config"
	"github.com/farhan-helmy/sedekahje-be/internal/db"
)

func main() {

	cfg := config.LoadConfig()

	mongoClient := db.ConnectDB(cfg.MongoURI)

	// Initialize Fiber app
	app := fiber.New()

}
