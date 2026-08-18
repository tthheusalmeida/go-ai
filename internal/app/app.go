package app

import (
	"net/http"

	"go-ai/internal/ai"
	"go-ai/internal/handler"
	"go-ai/internal/repository"
	"go-ai/internal/router"
	"go-ai/internal/service"
)

func newHealthHandler() *handler.HealthHandler {
	repository := repository.NewInMemoryHealthRepository()
	service := service.NewHealthService(repository)

	return handler.NewHealthHandler(service)
}

func newAIService(provider ai.Provider) *service.AIService {
	return service.NewAIService(provider)
}

func New(provider ai.Provider) http.Handler {
	healthHandler := newHealthHandler()

  aiService := service.NewAIService(provider)
  aiHandler := handler.NewAIHandler(aiService)

	return router.New(
		healthHandler,
    aiHandler,
	)
}
