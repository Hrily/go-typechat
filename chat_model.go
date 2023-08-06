package typechat

import (
	"context"
	"fmt"
	"os"

	"github.com/pkg/errors"
	openai "github.com/sashabaranov/go-openai"
)

type ChatModel interface {
	Send(ctx context.Context, system string, messages []string) (string, error)
}

type openAIChatModel struct {
	client *openai.Client
	model  string
}

func NewOpenAIChatModel() ChatModel {
	model := openai.GPT3Dot5Turbo
	if modelEnv := os.Getenv("OPENAI_MODEL"); modelEnv != "" {
		model = modelEnv
	}

	client := NewOpenAIClient()
	return &openAIChatModel{
		client: client,
		model:  model,
	}
}

func (m *openAIChatModel) Send(
	ctx context.Context, system string, messages []string,
) (string, error) {
	msgs := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: system,
		},
	}
	for _, message := range messages {
		msgs = append(msgs, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: message,
		})
	}

	resp, err := m.client.CreateChatCompletion(
		ctx, openai.ChatCompletionRequest{
			Model:    openai.GPT3Dot5Turbo,
			Messages: msgs,
		},
	)

	if err != nil {
		return "", errors.Wrap(err, "failed to send messages")
	}
	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("no choices returned")
	}

	return resp.Choices[0].Message.Content, nil
}
