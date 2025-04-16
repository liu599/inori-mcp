package tests

import (
	"context"
	"testing"

	InoriMCP "github.com/liu599/inori-mcp"
	"github.com/liu599/inori-mcp/tools"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/stretchr/testify/assert"
)

// 定义请求参数结构体
type HelloWorldParams struct {
	Name string `json:"name"`
}

func TestHelloWorld(t *testing.T) {
	// 初始化测试配置
	baseURL := "https://api.openai.com/v1"
	apiKey := "test-api-key"
	tools.InitLLMClient(baseURL, apiKey)

	// 测试用例
	tests := []struct {
		name          string
		params        HelloWorldParams
		expectedError bool
	}{
		{
			name: "正常调用",
			params: HelloWorldParams{
				Name: "小明",
			},
			expectedError: false,
		},
		{
			name:          "缺少name参数",
			params:        HelloWorldParams{},
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 创建上下文
			ctx := context.Background()

			// 创建请求
			request := mcp.CallToolRequest{
				Params: mcp.CallToolParams{
					Arguments: map[string]interface{}{
						"name": tt.params.Name,
					},
				},
			}

			// 调用处理函数
			result, err := tools.BookSearchTool.Handler(ctx, request)

			// 验证结果
			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.NotEmpty(t, result.Text)
			}
		})
	}
}

func TestHelloWorldRegistration(t *testing.T) {
	// 创建 MCP 服务器
	mcpServer := server.NewMCPServer(":8080", nil)

	// 注册工具
	tools.AddBookSearchTools(mcpServer)

	// 验证工具是否已注册
	tools := mcpServer.GetTools()
	assert.NotEmpty(t, tools)
	assert.Equal(t, "hello_world", tools[0].Name)
	assert.Equal(t, "使用大模型生成创意问候语", tools[0].Description)
}
