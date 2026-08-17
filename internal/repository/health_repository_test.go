package repository

import "testing"

func TestGetHealth(t *testing.T) {
	repository := NewInMemoryHealthRepository()

	health := repository.GetHealth()

	if health == nil {
		t.Fatal("expected health, got nil")
	}

	if health.Status != "ok" {
		t.Errorf("expected status ok, got %s", health.Status)
	}
}
