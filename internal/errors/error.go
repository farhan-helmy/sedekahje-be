package errors

import (
	"fmt"
	"net/http"
)

type CustomError struct {
	OriginalErr error  `json:"-"`
	Message     string `json:"message,omitempty"`
	StatusCode  int    `json:"status,omitempty"`
}

// Error returns the error message
func (e *CustomError) Error() string {
	if e.OriginalErr != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.OriginalErr)
	}
	return e.Message
}

// Unwrap returns the original error
func (e *CustomError) Unwrap() error {
	return e.OriginalErr
}

// New creates a new CustomError with an optional original error
func New(message string, statusCode int, originalErr error) *CustomError {
	return &CustomError{
		Message:     message,
		StatusCode:  statusCode,
		OriginalErr: originalErr,
	}
}

func Render(w http.ResponseWriter, err error, originalErr error) {
	c, ok := err.(*CustomError)

	if ok {
		var errorMessage string
		if originalErr != nil {
			errorMessage = fmt.Sprintf("%s - %v", c.Message, originalErr)
		} else {
			errorMessage = c.Message
		}
		http.Error(w, errorMessage, c.StatusCode)
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
