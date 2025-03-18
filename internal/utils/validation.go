package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type ValidationResponse struct {
	Status string            `json:"status"`
	Errors map[string]string `json:"errors"`
}

// FormatValidationErrors formats validator.ValidationErrors into a readable format
func FormatValidationErrors(err error) ValidationResponse {
	validationErrors := make(map[string]string)

	for _, err := range err.(validator.ValidationErrors) {
		field := err.Field()
		validationErrors[field] = fmt.Sprintf("Field '%s' validation failed: %s", field, err.Tag())
	}

	return ValidationResponse{
		Status: "error",
		Errors: validationErrors,
	}
}

// RespondWithValidationErrors writes validation errors to the http.ResponseWriter
func RespondWithValidationErrors(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(FormatValidationErrors(err))
}
