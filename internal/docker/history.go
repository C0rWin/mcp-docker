package docker

import (
	"context"
	"fmt"
	"os/exec"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// HistoryTool is the tool definition for the Docker history command.
var HistoryTool = mcp.NewTool("docker_history",
	mcp.WithDescription("Shows the history of an image"),
	mcp.WithString("image",
		mcp.Required(),
		mcp.Description("The name of the image to show history for"),
	),
	mcp.WithString("format",
		mcp.Description("Format the output using a custom template"),
	),
	mcp.WithString("no-trunc",
		mcp.Description("Don't truncate output"),
	),
	mcp.WithString("human",
		mcp.Description("Format the output in human-readable format"),
	),
)

// HistoryHandler is the handler function that handles history requests
func HistoryHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	image := req.Params.Arguments["image"].(string)

	args := []string{"history", image}

	format, exist := req.Params.Arguments["format"]
	if exist {
		args = append(args, "--format", format.(string))
	}

	_, exist = req.Params.Arguments["no-trunc"]
	if exist {
		args = append(args, "--no-trunc")
	}

	_, exist = req.Params.Arguments["human"]
	if exist {
		args = append(args, "--human")
	}

	historyCmd := exec.Command("docker", args...)
	outBytes, err := historyCmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to show history for image %s: %w", image, err)
	}
	result := string(outBytes)

	return mcp.NewToolResultText(result), nil
}

func WithHistoryTool(s *server.MCPServer) *server.MCPServer {
	s.AddTool(HistoryTool, HistoryHandler)
	return s
}
