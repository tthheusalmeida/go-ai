package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"go-ai/internal/model"
	"go-ai/internal/service"
)

type fakeHealthRepository struct {
	health *model.Health
}

func (f *fakeHealthRepository) GetHealth() *model.Health {
	return f.health
}

func TestGetHealth(t *testing.T) {
	expected := &model.Health{
		Status: "ok",
	}

	repository := &fakeHealthRepository{
		health: expected,
	}

	healthService := service.NewHealthService(repository)

	handler := NewHealthHandler(healthService)

	request := httptest.NewRequest("GET", "/health", nil)
	response := httptest.NewRecorder()

	handler.GetHealth(response, request)

	if response.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", response.Code)
	}

	if response.Header().Get("Content-Type") != "application/json" {
		t.Errorf("expected application/json, got %s",
			response.Header().Get("Content-Type"))
	}

  var result model.Health

  err := json.NewDecoder(response.Body).Decode(&result)
  if err != nil {
    t.Fatalf("failed to decode response: %v", err)
  }

  if result.Status != "ok" {
    t.Errorf("expected status ok, got %s", result.Status)
  }
}
