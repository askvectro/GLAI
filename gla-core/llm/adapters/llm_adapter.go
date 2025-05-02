package adapters

import (
	"bytes"
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"os/exec"

	openai "github.com/sashabaranov/go-openai"
)

// LLMAdapter defines a common interface for all LLM providers
type LLMAdapter interface {
	Generate(prompt string) (string, error)
}

// -------- OpenAI Adapter (GPT-4 / GPT-3.5) --------

type OpenAIAdapter struct {
	client *openai.Client
	model  string
}

func NewOpenAIAdapter(apiKey, model string) *OpenAIAdapter {
	client := openai.NewClient(apiKey)
	return &OpenAIAdapter{
		client: client,
		model:  model,
	}
}

func (oa *OpenAIAdapter) Generate(prompt string) (string, error) {
	resp, err := oa.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: oa.model,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)
	if err != nil {
		return "", err
	}
	return resp.Choices[0].Message.Content, nil
}

// -------- CodeLlama Adapter (Local CLI or subprocess) --------

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

// -------- CoreML Adapter (REST bridge from iOS) --------

type CoreMLAdapter struct {
	endpoint string
}

func NewCoreMLAdapter(endpoint string) *CoreMLAdapter {
	return &CoreMLAdapter{endpoint: endpoint}
}

func (c *CoreMLAdapter) Generate(prompt string) (string, error) {
	resp, err := http.Post(c.endpoint, "application/json", bytes.NewBuffer([]byte(`{"prompt":"`+prompt+`"}`)))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

// -------- Factory for Adapter Initialization --------

func LoadAdapter(name string, config map[string]string) (LLMAdapter, error) {
	switch name {
	case "openai":
		return NewOpenAIAdapter(config["api_key"], config["model"]), nil
	case "codellama":
		return NewCodeLlamaAdapter(config["bin_path"]), nil
	case "coreml":
		return NewCoreMLAdapter(config["endpoint"]), nil
	default:
		return nil, errors.New("unknown adapter type")
	}
}
