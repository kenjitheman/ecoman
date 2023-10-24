package openai

import (
	"testing"
)

func TestGenerateAdvice(t *testing.T) {
	result := "hello"
	advice := GenerateAdvice(result)
	if advice == "" {
		t.Errorf("GenerateAdvice() failed, got: %s, want: %s.", advice, result)
	}
}
