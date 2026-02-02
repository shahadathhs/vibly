package utils

import (
	"encoding/json"
	"net/http"

	"vibly/app/models"
)

func JSONResponse(w http.ResponseWriter, status int, success bool, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	response := models.APIResponse{
		Success: success,
		Message: message,
		Data:    data,
	}

	json.NewEncoder(w).Encode(response)
}

func ErrorResponse(w http.ResponseWriter, status int, err error) {
	var errorMessage string

	switch {
	case err != nil:
		errorMessage = err.Error()
	case status >= 500:
		errorMessage = "Internal server error"
	case status == http.StatusUnauthorized:
		errorMessage = "Unauthorized"
	case status == http.StatusBadRequest:
		errorMessage = "Bad request"
	default:
		errorMessage = "An error occurred"
	}

	JSONResponse(w, status, false, "", map[string]string{"error": errorMessage})
}
