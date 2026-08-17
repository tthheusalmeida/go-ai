package handler

import (
	"encoding/json"
	"net/http"

	"go-ai/internal/service"
)

type HealthHandler struct {
	service *service.HealthService
}

func NewHealthHandler(service *service.HealthService) *HealthHandler {
	return &HealthHandler{
		service: service,
	}
}

func (h *HealthHandler) GetHealth(w http.ResponseWriter, r *http.Request) {
	health := h.service.GetHealth()

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(health)
}