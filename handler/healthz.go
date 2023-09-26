package handler

import (
	"net/http"
	"encoding/json"
	"log"

	"github.com/TechBowl-japan/go-stations/model"
)

// A HealthzHandler implements health check endpoint.
type HealthzHandler struct{
	Message string `json:"message"`
}

// NewHealthzHandler returns HealthzHandler based http.Handler.
func NewHealthzHandler() *HealthzHandler {
	return &HealthzHandler{}
}

// ServeHTTP implements http.Handler interface.
func (h *HealthzHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(&model.HealthzResponse{
		Message: "OK",
	})
	if err != nil {
		log.Println(err)
	}
}
