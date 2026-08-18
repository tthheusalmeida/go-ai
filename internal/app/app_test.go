package app

import (
	"go-ai/internal/ai"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNew(t *testing.T) {
	provider := &ai.FakeProvider{
		Response: "test response",
	}

	router := New(provider)

	if router == nil {
		t.Fatal("expected application")
	}

	request := httptest.NewRequest(
		http.MethodGet,
		"/health",
		nil,
	)

	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Errorf(
			"expected status 200, got %d",
			response.Code,
		)
	}
}
