package docker

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

var PullTool = mcp.NewTool("docker_pull",
	mcp.WithDescription("Pulls a Docker image from a registry"),
	mcp.WithString("image",
		mcp.Required(),
		mcp.Description("The name of the image to pull"),
	),
)

// PullHandler is the handler function that handles pull requests
func PullHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	image := req.Params.Arguments["image"].(string)

	var stderr bytes.Buffer
	pullCmd := exec.Command("docker", "pull", image)
	pullCmd.Stderr = &stderr
	outBytes, err := pullCmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to pull image: %w \n %s", err, stderr.String())
	}
	result := string(outBytes)

	return mcp.NewToolResultText(fmt.Sprintf("%s", result)), nil
}

// WithPullTool adds the pull tool to the MCP server
func WithPullTool(s *server.MCPServer) *server.MCPServer {
	s.AddTool(PullTool, PullHandler)
	return s
}
