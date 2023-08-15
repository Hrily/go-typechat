package typechat

import (
	"os"

	openai "github.com/sashabaranov/go-openai"
)

// NewOpenAIClient ...
func NewOpenAIClient() *openai.Client {
	apiKey := os.Getenv("OPENAI_API_KEY")
	config := openai.DefaultConfig(apiKey)
	if orgID := os.Getenv("OPENAI_ORGANIZATION"); orgID != "" {
		config.OrgID = orgID
	}
	client := openai.NewClientWithConfig(config)

	return client
}
