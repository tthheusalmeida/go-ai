package config

import (
	"log"
	"os"
)

type Env struct {
  Port string
  Host string

  AIProvider string
  AIBaseURL string
  AIModel string
}

func Load() Env {
  return Env{
    Port: GetEnv("PORT", "8080"),
    Host: GetEnv("HOST", "localhost"),

    AIProvider: GetEnv("AI_PROVIDER", "ollama"),
    AIBaseURL:  GetEnv("AI_BASE_URL", "http://localhost:11434"),
    AIModel:    GetEnv("AI_MODEL", "qwen3:4b"),
  }
}

func GetEnv(env string, fallback string) string {
  value, exists := os.LookupEnv(env)

  if exists {
    return value
  }

  if fallback == "" {
    log.Fatalf("Variável de ambiente %q não está definida e não possui alternativa de fallback.", env)
  }

  return fallback
}
