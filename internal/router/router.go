package router

import (
	"net/http"

	"go-ai/internal/handler"
)

func New(handler *handler.HealthHandler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", handler.GetHealth)

	return mux
}