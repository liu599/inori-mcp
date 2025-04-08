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
	// 如果 baseURL 和 apiKey 都为空，尝试从配置文件加载
	if baseURL == "" && apiKey == "" {
		config, err := LoadConfig("config.prod.yaml")
		if err != nil {
			fmt.Println("请确保config.yaml已重命名为config.prod.yaml")
		} else {
			baseURL, apiKey = config.GetLLMConfig()
		}
	}

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
			Model: "deepseek-r1",
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
