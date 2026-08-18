package ai

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestOllamaProvider_Generate(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}

		if r.URL.Path != "/api/generate" {
			t.Errorf("expected /api/generate, got %s", r.URL.Path)
		}

		var request ollamaRequest

		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			t.Fatalf("failed to decode request: %v", err)
		}

		if request.Model != "test-model" {
			t.Errorf("expected model test-model, got %s", request.Model)
		}

		if request.Prompt != "Olá" {
			t.Errorf("expected prompt Olá, got %s", request.Prompt)
		}

		if request.Stream {
			t.Errorf("expected stream to be false")
		}

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(ollamaResponse{
			Response: "Olá! Como posso ajudar?",
		})
	}))
	defer server.Close()

	provider := NewOllamaProvider(
		server.URL,
		"test-model",
	)

	response, err := provider.Generate("Olá")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if response != "Olá! Como posso ajudar?" {
		t.Errorf("expected response, got %s", response)
	}
}

func TestOllamaProvider_Generate_HTTPError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}))
	defer server.Close()

	provider := NewOllamaProvider(
		server.URL,
		"test-model",
	)

	_, err := provider.Generate("Olá")

	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestOllamaProvider_Generate_InvalidJSON(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`invalid json`))
	}))
	defer server.Close()

	provider := NewOllamaProvider(
		server.URL,
		"test-model",
	)

	_, err := provider.Generate("Olá")

	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
