package config

import (
	"os"
	"os/exec"
	"testing"
)

func TestGetEnv(t *testing.T) {
  t.Setenv("PORT", "9000")

  value := GetEnv("PORT", "8080")

  if value != "9000" {
    t.Errorf("expected 9000, got %s", value)
  }
}

func TestGetEnvFallback(t *testing.T) {
  value := GetEnv("PORT_NOT_EXISTS", "8080")

	if value != "8080" {
		t.Errorf("expected 8080, got %s", value)
	}
}

func TestGetEnvWithoutFallback(t *testing.T) {
  if os.Getenv("GO_TEST_GET_ENV") == "1" {
    GetEnv("PORT_NOT_EXISTS", "")
    return
  }

  cmd := exec.Command(os.Args[0], "-test.run=TestGetEnvWithoutFallback")
  cmd.Env = append(os.Environ(), "GO_TEST_GET_ENV=1")

  err := cmd.Run()

  if err == nil {
    t.Fatal("expected GetEnv to terminate the process")
  }

  if exitError, ok := err.(*exec.ExitError); ok {
		if exitError.ExitCode() == 0 {
			t.Fatal("expected non-zero exit code")
		}
		return
	}

	t.Fatalf("unexpected error: %v", err)
}

func TestLoad(t *testing.T) {
	t.Setenv("PORT", "9000")
	t.Setenv("HOST", "localhost")

	env := Load()

	if env.Port != "9000" {
		t.Errorf("expected port 9000, got %s", env.Port)
	}

	if env.Host != "localhost" {
		t.Errorf("expected host localhost, got %s", env.Host)
	}
}
