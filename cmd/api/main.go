package main

import (
	"log"
	"net/http"

	"go-ai/config"
	"go-ai/internal/ai"
	"go-ai/internal/app"
)

func main() {
  env := config.Load()
  provider, err := ai.NewProvider(
    env.AIProvider,
    env.AIBaseURL,
    env.AIModel,
  )
  if err != nil {
    log.Fatal(err)
  }

  router := app.New(provider)

	server := &http.Server{
		Addr:    ":"+env.Port,
		Handler: router,
	}

	log.Printf("[GO-AI] running on http://%s:%s\n", env.Host, env.Port)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
