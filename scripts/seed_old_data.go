package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/farhan-helmy/sedekahje-be/internal/db"
	"github.com/farhan-helmy/sedekahje-be/internal/models"
	"github.com/farhan-helmy/sedekahje-be/internal/utils"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env Doesnt exist")
	}

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Println("MONGO_URI is not set")
	}

	client := db.ConnectDB(mongoURI)

	jsonFile, err := os.ReadFile("data/sedekahjeData.json")
	if err != nil {
		log.Println("Error opening file", err)
	}

	var institutions []models.Institution
	err = json.Unmarshal(jsonFile, &institutions)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	collection := client.Database("sedekahje").Collection("institutions")

	if _, err := collection.DeleteMany(context.Background(), bson.M{}); err != nil {
		log.Fatal("Error during DeleteMany(): ", err)
	}

	for _, v := range institutions {
		v.Slug = utils.Slugify(v.Name)

		_, err := collection.InsertOne(context.Background(), v)
		if err != nil {
			log.Fatal("Error during InsertOne(): ", err)
		}

		log.Printf("Inserted institution: %v", v.Name)
	}

	log.Println("Done")

	db.DisconnectDB(client)

}
