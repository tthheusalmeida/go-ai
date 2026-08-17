package config

import (
	"log"
	"os"
)

type Env struct {
  Port string
  Host string
}

func Load() Env {
  return Env{
    Port: GetEnv("PORT", "8080"),
    Host: GetEnv("HOST", "localhost"),
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
