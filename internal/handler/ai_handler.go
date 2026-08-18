package handler

import (
	"encoding/json"
	"net/http"

	"go-ai/internal/service"
)

type AIHandler struct {
  service *service.AIService
}

type generateRequest struct {
	Prompt string `json:"prompt"`
}

type generateResponse struct {
	Response string `json:"response"`
}

func NewAIHandler(service *service.AIService) *AIHandler {
  return &AIHandler{
    service: service,
  }
}

func (h *AIHandler) Generate(w http.ResponseWriter, r *http.Request) {
  var request generateRequest

  err := json.NewDecoder(r.Body).Decode(&request)
  if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

  response, err := h.service.Generate(request.Prompt)
	if err != nil {
		http.Error(w, "AI failed to generate response", http.StatusInternalServerError)
		return
	}

  w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(generateResponse{
		Response: response,
	})
}
