package InoriMCP

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

// LLMClient 大模型客户端配置
type LLMClient struct {
	Client *openai.Client
}

// NewLLMClient 创建新的LLM客户端
func NewLLMClient(baseURL, apiKey string) *LLMClient {
	config := openai.DefaultConfig(apiKey)
	if baseURL != "" {
		config.BaseURL = baseURL
	}

	return &LLMClient{
		Client: openai.NewClientWithConfig(config),
	}
}

// GenerateText 生成文本
func (c *LLMClient) GenerateText(ctx context.Context, prompt string) (string, error) {
	resp, err := c.Client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)

	if err != nil {
		return "", fmt.Errorf("LLM生成失败: %v", err)
	}

	return resp.Choices[0].Message.Content, nil
} 