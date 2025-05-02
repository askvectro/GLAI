package pipeline

import (
	"go/format"
)

type Formatter interface {
	Format(code string) string
}

// GoFormatter formats Go code using go/format
type GoFormatter struct{}

func NewGoFormatter() *GoFormatter {
	return &GoFormatter{}
}

func (gf *GoFormatter) Format(code string) string {
	formatted, err := format.Source([]byte(code))
	if err != nil {
		// Return unformatted if formatting fails
		return code
	}
	return string(formatted)
}
