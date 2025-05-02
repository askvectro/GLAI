package tests

import (
	"testing"
	"gla-core/llm/adapters"
)

func TestCodeLlamaAdapter_Generate(t *testing.T) {
	adapter := adapters.NewCodeLlamaAdapter("echo") // mock binary
	output, err := adapter.Generate("hello")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if output == "" {
		t.Error("expected non-empty output")
	}
}

func TestCoreMLAdapter_Generate_InvalidURL(t *testing.T) {
	adapter := adapters.NewCoreMLAdapter("http://invalid-url")
	_, err := adapter.Generate("test prompt")
	if err == nil {
		t.Error("expected error for invalid URL, got nil")
	}
}
