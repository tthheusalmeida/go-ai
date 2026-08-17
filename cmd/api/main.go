package main

import (
	"log"
	"net/http"

	"go-ai/internal/handler"
	"go-ai/internal/repository"
	"go-ai/internal/router"
	"go-ai/internal/service"
)

func main() {
	healthRepository := repository.NewInMemoryHealthRepository()

	healthService := service.NewHealthService(
		healthRepository,
	)

	healthHandler := handler.NewHealthHandler(
		healthService,
	)

	router := router.New(
		healthHandler,
	)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Println("[GO-AI] running on http://localhost:8080")

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
