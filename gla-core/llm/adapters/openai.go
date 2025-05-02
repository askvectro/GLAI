package adapters

import (
	"context"
	"fmt"
	openai "github.com/sashabaranov/go-openai"
)

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
