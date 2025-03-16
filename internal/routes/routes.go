package routes

import (
	"encoding/json"
	"net/http"

	"github.com/farhan-helmy/sedekahje-be/internal/models"
	"github.com/farhan-helmy/sedekahje-be/internal/services"
	"github.com/farhan-helmy/sedekahje-be/internal/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var validate *validator.Validate

func init() {
	validate = validator.New() // Initialize validator once
}

func SetupRoutes(router *mux.Router, client *mongo.Client) {
	institutionService := services.NewInstitutionService(client)

	api := router.PathPrefix("/api").Subrouter()

	api.HandleFunc("/health", utils.HealthCheckHandler).Methods("GET")

	api.HandleFunc("/institutions", getAllInstitutions(institutionService)).Methods("GET")
	api.HandleFunc("/institutions", createInstitution(institutionService)).Methods("POST")
}

func getAllInstitutions(institutionService *services.InstitutionService) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		institutions, err := institutionService.GetInstitutions()

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(institutions)
	}
}

func createInstitution(institutionService *services.InstitutionService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var institution models.Institution

		if err := json.NewDecoder(r.Body).Decode(&institution); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if err := validate.Struct(institution); err != nil {
			// Return validation errors
			errs := err.(validator.ValidationErrors)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": errs.Error()})
			return
		}

		json.NewEncoder(w).Encode(institution)
	}
}
