package ai

import "testing"

func TestNewProviderFake(t *testing.T) {
	provider, err := NewProvider("fake", "", "")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if provider == nil {
		t.Fatal("expected provider")
	}
}

func TestNewProviderOllama(t *testing.T) {
	provider, err := NewProvider(
		"ollama",
		"http://localhost:11434",
		"gemma3",
	)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	ollamaProvider, ok := provider.(*OllamaProvider)
	if !ok {
		t.Fatalf("expected *OllamaProvider, got %T", provider)
	}

	if ollamaProvider.baseURL != "http://localhost:11434" {
		t.Errorf(
			"expected baseURL http://localhost:11434, got %s",
			ollamaProvider.baseURL,
		)
	}

	if ollamaProvider.model != "gemma3" {
		t.Errorf(
			"expected model gemma3, got %s",
			ollamaProvider.model,
		)
	}
}

func TestNewProviderUnsupported(t *testing.T) {
	provider, err := NewProvider("unknown", "", "")

	if err == nil {
		t.Fatal("expected error")
	}

	if provider != nil {
		t.Fatal("expected nil provider")
	}
}
