package service

import (
	"testing"

	"go-ai/internal/model"
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

	service := NewHealthService(repository)

	result := service.GetHealth()

	if result != expected {
		t.Fatal("expected service to return repository health")
	}
}
