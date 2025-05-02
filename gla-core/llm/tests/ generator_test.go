package tests

import (
	"testing"
	"gla-core/llm/pipeline"
)

// MockAdapter simulates LLMAdapter for tests
type MockAdapter struct {
	Output string
	Err    error
}

func (m *MockAdapter) Generate(prompt string) (string, error) {
	return m.Output, m.Err
}

// MockFormatter formats by adding a prefix
type MockFormatter struct{}

func (mf *MockFormatter) Format(code string) string {
	return "// formatted\n" + code
}

func TestGenerator_Run_Success(t *testing.T) {
	adapter := &MockAdapter{Output: "func main() {}"}
	filter := pipeline.NewSimpleFilter([]string{"panic"})
	formatter := &MockFormatter{}

	gen := pipeline.NewGenerator(adapter, []pipeline.Filter{filter}, formatter)

	result, err := gen.Run("write a main function")
	if err != nil {
		t.Errorf("expected success, got error: %v", err)
	}

	if result == "" {
		t.Error("expected formatted output, got empty string")
	}

	if result[:13] != "// formatted\n" {
		t.Errorf("formatter not applied, got: %s", result)
	}
}

func TestGenerator_Run_FilterRejects(t *testing.T) {
	adapter := &MockAdapter{Output: "panic(\"unsafe\")"}
	filter := pipeline.NewSimpleFilter([]string{"panic"})
	formatter := &MockFormatter{}

	gen := pipeline.NewGenerator(adapter, []pipeline.Filter{filter}, formatter)

	_, err := gen.Run("generate panic code")
	if err == nil {
		t.Error("expected filter rejection error, got nil")
	}
}
