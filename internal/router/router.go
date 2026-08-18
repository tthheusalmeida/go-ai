package router

import (
	"net/http"

	"go-ai/internal/handler"
)

func New(
  healthHandler *handler.HealthHandler,
	aiHandler *handler.AIHandler,
) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", healthHandler.GetHealth)
	mux.HandleFunc("POST /ai/generate", aiHandler.Generate)

	return mux
}
