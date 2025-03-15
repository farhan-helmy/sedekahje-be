package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoURI string
}

func LoadConfig() Config {
	godotenv.Load()

	return Config{
		MongoURI: os.Getenv("MONGO_URI"),
	}
}
