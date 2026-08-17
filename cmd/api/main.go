package main

import (
	"log"
	"net/http"

	"go-ai/config"
	"go-ai/internal/app"
)

func main() {
  env := config.Load()
  router := app.New()

	server := &http.Server{
		Addr:    ":"+env.Port,
		Handler: router,
	}

	log.Printf("[GO-AI] running on http://%s:%s\n", env.Host, env.Port)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
