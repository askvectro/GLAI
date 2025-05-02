package pipeline

import "strings"

type Filter interface {
	Approve(output string) bool
	Name() string
}

// SimpleFilter checks for banned words (basic safety check)
type SimpleFilter struct {
	BannedWords []string
}

func NewSimpleFilter(banned []string) *SimpleFilter {
	return &SimpleFilter{BannedWords: banned}
}

func (sf *SimpleFilter) Approve(output string) bool {
	for _, word := range sf.BannedWords {
		if strings.Contains(strings.ToLower(output), strings.ToLower(word)) {
			return false
		}
	}
	return true
}

func (sf *SimpleFilter) Name() string {
	return "SimpleFilter"
}
