package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/leanderp/golang_rest_web/server"
)

type HomeResponse struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

func HomeHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response := HomeResponse{
			Message: "Welcome to the Golang REST Web API",
			Status:  true,
		}

		json.NewEncoder(w).Encode(response)

	}
}
