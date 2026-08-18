package ai

type FakeProvider struct {
  Response string
}

func (f *FakeProvider) Generate(prompt string) (string, error) {
  return f.Response, nil
}
