package ai

import "fmt"

func NewProvider(provider string, baseURL string, model string) (Provider, error) {
  switch provider {
  case "fake":
    return &FakeProvider{
      Response: "fake response",
    }, nil
  case "ollama":
    return NewOllamaProvider(baseURL, model), nil

  default:
    return nil, fmt.Errorf("unsupported AI provider: %s", provider)
  }
}
