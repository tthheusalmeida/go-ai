package router

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"go-ai/internal/ai"
	"go-ai/internal/handler"
	"go-ai/internal/model"
	"go-ai/internal/service"
)

type fakeHealthRepository struct {
	health *model.Health
}

func (f *fakeHealthRepository) GetHealth() *model.Health {
	return f.health
}

func setupRouter() http.Handler {
	repository := &fakeHealthRepository{
		health: &model.Health{
			Status: "ok",
		},
	}

	healthService := service.NewHealthService(repository)
	healthHandler := handler.NewHealthHandler(healthService)

	provider := &ai.FakeProvider{
		Response: "test response",
	}

	aiService := service.NewAIService(provider)
	aiHandler := handler.NewAIHandler(aiService)

	return New(healthHandler, aiHandler)
}

func TestHealthRoute(t *testing.T) {
	router := setupRouter()

	request := httptest.NewRequest(
		http.MethodGet,
		"/health",
		nil,
	)

	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", response.Code)
	}
}

func TestAIGenerateRoute(t *testing.T) {
	router := setupRouter()

	request := httptest.NewRequest(
		http.MethodPost,
		"/ai/generate",
		strings.NewReader(`{"prompt":"me responda"}`),
	)

	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", response.Code)
	}
}

func TestRouteNotFound(t *testing.T) {
	router := setupRouter()

	request := httptest.NewRequest(
		http.MethodGet,
		"/nao-existe",
		nil,
	)

	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	if response.Code != http.StatusNotFound {
		t.Errorf(
			"expected status 404, got %d",
			response.Code,
		)
	}
}
