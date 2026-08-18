package service

import "go-ai/internal/ai"

type AIService struct {
  provider ai.Provider
}

func NewAIService(provider ai.Provider) *AIService {
  return &AIService{
    provider: provider,
  }
}

func (s *AIService) Generate(prompt string) (string, error) {
  return s.provider.Generate(prompt)
}
