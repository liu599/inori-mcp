package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/liu599/inori-mcp/tools"
	"log"
	"os"

	"github.com/mark3labs/mcp-go/server"
)

func newServer() *server.MCPServer {
	s := server.NewMCPServer(
		"mcp-inori-server",
		"0.0.1",
	)
	tools.AddBookSearchTools(s)
	return s
}

func run(transport, addr string) error {
	s := newServer()

	switch transport {
	case "stdio":
		srv := server.NewStdioServer(s)
		log.Printf("Stdio server start")
		return srv.Listen(context.Background(), os.Stdin, os.Stdout)
	case "sse":
		url := "http://" + addr
		srv := server.NewSSEServer(s,
			server.WithBaseURL(url),
		)
		log.Printf("SSE server listening on %s", addr)
		if err := srv.Start(addr); err != nil {
			return fmt.Errorf("Server error: %v", err)
		}
	default:
		return fmt.Errorf(
			"Invalid transport type: %s. Must be 'stdio' or 'sse'",
			transport,
		)
	}
	return nil
}

func main() {
	var transport string
	flag.StringVar(&transport, "t", "stdio", "Transport type (stdio or sse)")
	flag.StringVar(
		&transport,
		"transport",
		"stdio",
		"Transport type (stdio or sse)",
	)
	addr := flag.String("sse-address", "localhost:8000", "The host and port to start the sse server on")
	flag.Parse()

	if err := run(transport, *addr); err != nil {
		panic(err)
	}
}
