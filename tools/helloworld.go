package tools

import (
	"context"
	"errors"
	"fmt"
	InoriMCP "github.com/liu599/inori-mcp"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// HelloWorld 处理函数
func helloWorldHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	name, ok := request.Params.Arguments["name"].(string)
	if !ok {
		return nil, errors.New("name must be a string")
	}

	prompt := fmt.Sprintf("请用创意的方式向%s说Hello", name)

	return mcp.NewToolResultText(prompt), nil
}

var HelloWorld = InoriMCP.MustTool(
	"hello_world",
	"使用大模型生成创意问候语",
	helloWorldHandler,
)

func AddHelloWorldTools(mcp *server.MCPServer) {
	HelloWorld.Register(mcp)
}
