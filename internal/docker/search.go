package docker

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

var SearchTool = mcp.NewTool("docker_search",
	mcp.WithDescription("Searches for Docker images"),
	mcp.WithString("query",
		mcp.Required(),
		mcp.Description("The search query"),
	),
	mcp.WithString("filter",
		mcp.Description("The filter to apply to the search"),
	),
	mcp.WithString("format",
		mcp.Description("The format to use for the output"),
	),
	mcp.WithString("limit",
		mcp.Description("The maximum number of results to return"),
	),
)

// SearchHandler is the handler function that handles search requests
func SearchHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	query := req.Params.Arguments["query"].(string)
	args := []string{"search", query}
	filteri, exist := req.Params.Arguments["filter"]
	if exist {
		args = append(args, "--filter", filteri.(string))
	}
	format, exist := req.Params.Arguments["format"]
	if exist {
		args = append(args, "--format", format.(string))
	}
	limit, exist := req.Params.Arguments["limit"]
	if exist {
		args = append(args, "--limit", limit.(string))
	}

	var stderr bytes.Buffer
	searchCmd := exec.Command("docker", args...)
	searchCmd.Stderr = &stderr
	outBytes, err := searchCmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to search for images: %w \n %s", err, stderr.String())
	}
	result := string(outBytes)

	return mcp.NewToolResultText(fmt.Sprintf("%s", result)), nil
}

// WithSearchTool adds the search tool to the MCP server
func WithSearchTool(s *server.MCPServer) *server.MCPServer {
	s.AddTool(SearchTool, SearchHandler)
	return s
}
