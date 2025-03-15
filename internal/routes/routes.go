package routes

import (
	"encoding/json"
	"net/http"

	"github.com/farhan-helmy/sedekahje-be/internal/models"
	"github.com/farhan-helmy/sedekahje-be/internal/services"
	"github.com/farhan-helmy/sedekahje-be/internal/utils"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func SetupRoutes(router *mux.Router, client *mongo.Client) {
	institutionService := services.NewInstitutionService(client)

	api := router.PathPrefix("/api").Subrouter()

	api.HandleFunc("/health", utils.HealthCheckHandler).Methods("GET")

	api.HandleFunc("/institutions", getAllInstitutions(institutionService)).Methods("GET")

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
		var institution models.Institution
		json.NewDecoder(r.Body).Decode(&institution)

		err := institutionService.CreateInstitution(institution)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		json.NewEncoder(w).Encode(institution)
	}
}
