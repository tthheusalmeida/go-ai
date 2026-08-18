package service

import (
	"testing"

	"go-ai/internal/ai"
)

func TestAIServiceGenerate(t *testing.T) {
	provider := &ai.FakeProvider{
		Response: "resposta da IA",
	}

	service := NewAIService(provider)

	result, err := service.Generate("me responda")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result != "resposta da IA" {
		t.Errorf("expected resposta da IA, got %s", result)
	}
}
