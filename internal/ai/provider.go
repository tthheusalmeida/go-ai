package ai

type Provider interface {
  Generate(prompt string) (string, error)
}
