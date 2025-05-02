package pipeline

import (
	"fmt"
	"gla-core/llm/adapters"
)

type Generator struct {
	Adapter adapters.LLMAdapter
	Filters []Filter
	Formatter Formatter
}

func NewGenerator(adapter adapters.LLMAdapter, filters []Filter, formatter Formatter) *Generator {
	return &Generator{
		Adapter: adapter,
		Filters: filters,
		Formatter: formatter,
	}
}

func (g *Generator) Run(prompt string) (string, error) {
	raw, err := g.Adapter.Generate(prompt)
	if err != nil {
		return "", fmt.Errorf("generation failed: %w", err)
	}

	for _, filter := range g.Filters {
		if !filter.Approve(raw) {
			return "", fmt.Errorf("content rejected by filter: %s", filter.Name())
		}
	}

	formatted := g.Formatter.Format(raw)
	return formatted, nil
}
