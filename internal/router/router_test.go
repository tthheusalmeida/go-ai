package router

import (
	"net/http"
	"net/http/httptest"
	"testing"

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

func TestHealthRoute(t *testing.T) {
	repository := &fakeHealthRepository{
		health: &model.Health{
			Status: "ok",
		},
	}

	healthService := service.NewHealthService(repository)
	healthHandler := handler.NewHealthHandler(healthService)

	router := New(healthHandler)

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

func TestRouteNotFound(t *testing.T) {
	repository := &fakeHealthRepository{
		health: &model.Health{
			Status: "ok",
		},
	}

	healthService := service.NewHealthService(repository)
	healthHandler := handler.NewHealthHandler(healthService)

	router := New(healthHandler)

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
