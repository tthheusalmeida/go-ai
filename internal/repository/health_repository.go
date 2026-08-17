package repository

import "go-ai/internal/model"

type HealthRepository interface {
	GetHealth() *model.Health
}

type InMemoryHealthRepository struct{}

func NewInMemoryHealthRepository() *InMemoryHealthRepository {
	return &InMemoryHealthRepository{}
}

func (r *InMemoryHealthRepository) GetHealth() *model.Health {
	return &model.Health{
		Status: "ok",
	}
}
