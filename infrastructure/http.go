package infrastructure

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func JsonResponse(w http.ResponseWriter, statusCode int, message string) error {
	w.Header().Set("Content-Type", "application/json")

	rsp := Response{Message: message}
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(rsp)
}

func JSON(w http.ResponseWriter, statusCode int, body interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(body)
}
