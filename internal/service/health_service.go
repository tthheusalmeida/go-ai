package service

import (
	"go-ai/internal/model"
	"go-ai/internal/repository"
)

type HealthService struct {
	repository repository.HealthRepository
}

func NewHealthService(repository repository.HealthRepository) *HealthService {
	return &HealthService{
		repository: repository,
	}
}

func (s *HealthService) GetHealth() *model.Health {
	return s.repository.GetHealth()
}