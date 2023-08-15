package typechat

import (
	"context"
	"fmt"
	"os"

	"github.com/pkg/errors"
	openai "github.com/sashabaranov/go-openai"
)

// ChatModel is an interface to chat based language model
type ChatModel interface {
	Send(ctx context.Context, messages []*ChatModelMessage) (string, error)
}

// ChatModelMessage is a message to send to the chat model
// Invariant: one of System, User, or AI must be non-nil
type ChatModelMessage struct {
	System *string
	User   *string
	AI     *string
}

type openAIChatModel struct {
	client *openai.Client
	model  string
}

// NewOpenAIChatModel ...
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

// Send messages to the chat model and return the response
func (m *openAIChatModel) Send(
	ctx context.Context, messages []*ChatModelMessage,
) (string, error) {
	msgs := make([]openai.ChatCompletionMessage, 0, len(messages))
	for _, message := range messages {
		openAIMessage, err := m.toOpenAIMessage(message)
		if err != nil {
			return "", errors.Wrap(err, "failed to create message")
		}
		msgs = append(msgs, *openAIMessage)
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

func (m *openAIChatModel) toOpenAIMessage(
	message *ChatModelMessage,
) (*openai.ChatCompletionMessage, error) {
	switch {
	case message.System != nil:
		return &openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleSystem,
			Content: *message.System,
		}, nil
	case message.User != nil:
		return &openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: *message.User,
		}, nil
	case message.AI != nil:
		return &openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleAssistant,
			Content: *message.AI,
		}, nil
	default:
		return nil, fmt.Errorf("invalid message")
	}
}
