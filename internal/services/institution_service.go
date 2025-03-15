package services

import (
	"context"

	"github.com/farhan-helmy/sedekahje-be/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type InstitutionService struct {
	collection *mongo.Collection
}

func NewInstitutionService(client *mongo.Client) *InstitutionService {
	collection := client.Database("sedekahje").Collection("institutions")
	return &InstitutionService{collection}
}

func (s *InstitutionService) CreateInstitution(institution models.Institution) error {
	_, err := s.collection.InsertOne(context.Background(), institution)

	return err
}

func (s *InstitutionService) GetInstitutions() ([]models.Institution, error) {
	var institutions []models.Institution

	cursor, err := s.collection.Find(context.Background(), bson.M{})

	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.Background())

	if err := cursor.All(context.Background(), &institutions); err != nil {
		return nil, err
	}

	return institutions, nil
}
