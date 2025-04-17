# A very simple MCP Server

本项目分别使用Go和Python实现了一个最简单的MCP Server, 并使用一个mcp client去调用。本项目不依赖任何外部API, 你可以很方便的理解一个简单的MCP工作流。

```
情景假设: 书架编号A,B,C,D有4架子书, 我需要大模型帮我查对应编号的书
```

-  [bookSearch.go](tools%2FbookSearch.go) MCP 工具的Go实现,
-  [MCPServer_book_search.py](client%2FMCPServer_book_search.py)  MCP工具的python实现
-  [README.md](client%2FREADME.md) python工作流

## 编写方法:

```
规定MCP的入参, json参数名字起的尽量容易被大模型理解。

type BookShelfParams struct {
	BookShelfId string `json:"book_shelf_id" jsonschema:"required,description=The ID of the book shelf"`
}

工具核心逻辑, 这里可以替换成任意外部API, 最后Stdio需要返回一个字符串。

func bookSearchHandler(ctx context.Context, args BookShelfParams) (*mcp.CallToolResult, error) {
    // 逻辑
    return mcp.NewToolResultText("返回结果:xxxx"), nil
}

编写该MCP的名称和说明, 尽量容易人类理解的语言

var BookSearchTool = InoriMCP.MustTool(
	"book-search",
	"图书搜索工具",
	bookSearchHandler,
)

最后写入口函数

func AddBookSearchTools(mcp *server.MCPServer) {
	BookSearchTool.Register(mcp)
}

在main.go中, 加入这个工具

tools.AddBookSearchTools(s)

build:  go build -o inori-mcp.exe ./cmd/inori-mcp

mcp的启动方式: mcp.exe

python-mcp和go-mcp只差启动方式。

```

## 调试

安装node环境, 使用inspector, 先Connect然后就可以输入A,B,C,D看看是否返回书架里的书

- `npx -y @modelcontextprotocol/inspector .\inori-mcp.exe`

## 启动Client与MCP联动

- `python client/client.py`

只需要更改[MCPClient.py](client%2FMCPClient.py)中51行的connect_to_server的启动脚本即可。

```
        server_params = StdioServerParameters(
            command=<mcp.exe>的路径,
            args=[],
            env=None
        )
```

需要从.env.sample中创建.env, 填入OpenAI Compatible大模型的base_url, 模型建议使用QWQ-32B(不建议DeepSeek系列)

模型需要触发ToolCall并总结, 最后返回书架的书。

