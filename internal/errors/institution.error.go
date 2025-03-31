package errors

import (
	"net/http"
)

var (
	ErrInstitutionNotFound = &CustomError{
		Message:    "Institution not found",
		StatusCode: http.StatusNotFound,
	}

	ErrInstitutionAlreadyExists = &CustomError{
		Message:    "Institution already exists",
		StatusCode: http.StatusConflict,
	}

	ErrEmptyInstitution = &CustomError{
		Message:    "Institution is empty",
		StatusCode: http.StatusNotFound,
	}
)
