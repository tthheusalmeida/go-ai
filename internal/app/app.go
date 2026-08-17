package app

import (
	"net/http"

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

func New() http.Handler {
	healthHandler := newHealthHandler()

	return router.New(
		healthHandler,
	)
}
