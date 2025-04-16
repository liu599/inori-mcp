package tools

import (
	"context"
	"fmt"
	InoriMCP "github.com/liu599/inori-mcp"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"strings"
)

type BookShelfParams struct {
	BookShelfId string `json:"book_shelf_id" jsonschema:"required,description=The ID of the book shelf"`
}

// BookSearchTool 处理函数
func bookSearchHandler(ctx context.Context, args BookShelfParams) (*mcp.CallToolResult, error) {
	shelfID := args.BookShelfId

	var books []string
	switch shelfID {
	case "A":
		books = []string{
			"《三体》 - 作者：刘慈欣\n荣获雨果奖的科幻巨著，描绘了地球文明与三体文明的惊心动魄的第一次接触。",
			"《基地》 - 作者：艾萨克·阿西莫夫\n科幻文学的奠基之作，讲述了人类帝国衰落与重建的宏大故事。",
			"《神经漫游者》 - 作者：威廉·吉布森\n赛博朋克开山之作，描绘了一个由计算机网络主导的未来世界。",
			"《超新星纪元》 - 作者：刘慈欣\n一部描写儿童接管地球后发生的故事，展现了人性的复杂与深邃。",
		}
	case "B":
		books = []string{
			"《百年孤独》 - 作者：加西亚·马尔克斯\n魔幻现实主义代表作，讲述了布恩迪亚家族七代人的传奇故事。",
			"《霍乱时期的爱情》 - 作者：加西亚·马尔克斯\n一部跨越半个多世纪的爱情史诗，展现了爱情的永恒魅力。",
			"《悲惨世界》 - 作者：维克多·雨果\n描写了冉阿让等人物的救赎故事，展现了人性的光辉。",
			"《红楼梦》 - 作者：曹雪芹\n中国古典小说巅峰之作，描绘了大观园中的人生百态。",
		}
	case "C":
		books = []string{
			"《深度学习》 - 作者：Ian Goodfellow等\n人工智能领域的经典教材，全面介绍了深度学习的理论与实践。",
			"《算法导论》 - 作者：Thomas H. Cormen等\n计算机科学必读经典，系统讲解了各种算法的设计与分析。",
			"《计算机程序的构造和解释》 - 作者：Harold Abelson等\nMIT经典教材，探讨了程序设计的本质。",
			"《代码大全》 - 作者：Steve McConnell\n软件开发实践指南，涵盖了软件构建的各个方面。",
		}
	case "D":
		books = []string{
			"《魔法少女小圆》 - 作者：Magica Quartet\n颠覆传统魔法少女题材的作品，探讨了希望与绝望的永恒主题。",
			"《命运石之门》 - 作者：5pb.\n关于时间旅行的科幻故事，展现了蝴蝶效应带来的影响。",
			"《进击的巨人》 - 作者：谏山创\n人类与巨人的史诗级战斗，探讨了自由与压迫的主题。",
			"《钢之炼金术师》 - 作者：荒川弘\n讲述了寻找贤者之石的旅程，展现了等价交换的哲学。",
		}
	default:
		return nil, fmt.Errorf("未知的书架编号：%s，可用的书架编号为：A（科幻）、B（文学）、C（计算机）、D（动漫）", shelfID)
	}

	prompt := fmt.Sprintf("%s号书架的图书列表：\n\n%s", shelfID, strings.Join(books, "\n\n"))
	return mcp.NewToolResultText(prompt), nil
}

var BookSearchTool = InoriMCP.MustTool(
	"book-search",
	"图书搜索工具",
	bookSearchHandler,
)

func AddBookSearchTools(mcp *server.MCPServer) {
	BookSearchTool.Register(mcp)
}
