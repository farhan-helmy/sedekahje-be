package services

import (
	"context"
	"fmt"

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

func (s *InstitutionService) CreateInstitution(institution *models.Institution) error {
	if _, err := s.collection.InsertOne(context.Background(), institution); err != nil {
		return err
	}

	fmt.Println("Created institution", institution.Name)

	return nil
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

func (s *InstitutionService) GetInstitutionBySlug(slug string) (*models.Institution, error) {
	var institution models.Institution

	err := s.collection.FindOne(context.Background(), bson.M{"slug": slug}).Decode(&institution)

	return &institution, err
}
