package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"go-ai/internal/ai"
	"go-ai/internal/service"
)

func TestAIHandlerGenerate(t *testing.T) {
	provider := &ai.FakeProvider{
		Response: "resposta da IA",
	}

	aiService := service.NewAIService(provider)
	handler := NewAIHandler(aiService)

	request := httptest.NewRequest(
		http.MethodPost,
		"/ai/generate",
		strings.NewReader(`{"prompt":"me responda"}`),
	)

	response := httptest.NewRecorder()

	handler.Generate(response, request)

	if response.Code != http.StatusOK {
		t.Errorf(
			"expected status 200, got %d",
			response.Code,
		)
	}

	expected := `{"response":"resposta da IA"}`

	if strings.TrimSpace(response.Body.String()) != expected {
		t.Errorf(
			"expected %s, got %s",
			expected,
			response.Body.String(),
		)
	}
}
