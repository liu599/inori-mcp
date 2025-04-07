package tools

import (
	"context"
	"errors"
	"fmt"
	InoriMCP "github.com/liu599/inori-mcp"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

var llmClient *InoriMCP.LLMClient

// InitLLMClient 初始化LLM客户端
func InitLLMClient(baseURL, apiKey string) {
	llmClient = InoriMCP.NewLLMClient(baseURL, apiKey)
}

// HelloWorld 处理函数
func helloWorldHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	name, ok := request.Params.Arguments["name"].(string)
	if !ok {
		return nil, errors.New("name must be a string")
	}

	if llmClient == nil {
		return nil, errors.New("LLM客户端未初始化")
	}

	// 使用 LLM 生成问候语
	prompt := fmt.Sprintf("请用创意的方式向%s说Hello", name)
	response, err := llmClient.GenerateText(ctx, prompt)
	if err != nil {
		return nil, fmt.Errorf("LLM生成失败: %v", err)
	}

	return mcp.NewToolResultText(response), nil
}

var HelloWorld = InoriMCP.MustTool(
	"hello_world",
	"使用大模型生成创意问候语",
	helloWorldHandler,
)

func AddHelloWorldTools(mcp *server.MCPServer) {
	HelloWorld.Register(mcp)
}
