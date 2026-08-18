package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

var _ Provider = (*OllamaProvider)(nil)

type ollamaRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

type ollamaResponse struct {
	Response string `json:"response"`
}

type OllamaProvider struct {
	baseURL string
	model   string
}

func NewOllamaProvider(baseURL string, model string) *OllamaProvider {
	return &OllamaProvider{
		baseURL: baseURL,
		model:   model,
	}
}

func (p *OllamaProvider) Generate(prompt string) (string, error) {
  request := ollamaRequest{
    Model: p.model,
    Prompt: prompt,
    Stream: false,
  }

  body, err := json.Marshal(request)
  if err != nil {
    return "", err
  }

  response, err := http.Post(
    p.baseURL+"/api/generate",
    "application/json",
    bytes.NewBuffer(body),
  )
  if err != nil {
    return  "", err
  }

  defer response.Body.Close()

  if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("ollama returned status: %s", response.Status)
	}

	var result ollamaResponse

	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return "", err
	}

	return result.Response, nil
}
