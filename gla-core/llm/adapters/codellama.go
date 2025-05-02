package adapters

import (
	"bytes"
	"os/exec"
)

type CodeLlamaAdapter struct {
	binPath string
}

func NewCodeLlamaAdapter(binPath string) *CodeLlamaAdapter {
	return &CodeLlamaAdapter{binPath: binPath}
}

func (c *CodeLlamaAdapter) Generate(prompt string) (string, error) {
	cmd := exec.Command(c.binPath, "--prompt", prompt)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return out.String(), nil
}
