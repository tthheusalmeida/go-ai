package ai

import "testing"

func TestFakeProviderGenerate(t *testing.T) {
  fake := &FakeProvider{
    Response: "resposta da IA",
  }

  result, err := fake.Generate("me responda")

  if err != nil {
    t.Fatalf("unexpected error: %v", err)
  }

  if result != "resposta da IA" {
    t.Errorf("expected resposta da IA, got %s", result)
  }
}
