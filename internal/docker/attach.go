package docker

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

var AttachTool = mcp.NewTool("docker_attach",
	mcp.WithDescription("Attach to a running container"),
	mcp.WithString("containerID",
		mcp.Required(),
		mcp.Description("The ID of the container to attach to"),
	),
)

// AttachHandler is the handler function that handles attach requests
func AttachHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// Implement the logic to attach to a Docker container
	// This is a placeholder implementation
	containerID := req.Params.Arguments["containerID"].(string)

	var stderr bytes.Buffer
	diffCmd := exec.Command("docker", "attach", containerID)
	diffCmd.Stderr = &stderr

	outBytes, err := diffCmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to attach to container %s: %w\n %s", containerID, err, stderr.String())
	}
	result := string(outBytes)

	return mcp.NewToolResultText(fmt.Sprintf("%s", result)), nil
}

// WithAttachTool is a convenience function to add the AttachTool to the MCP server
func WithAttachTool(s *server.MCPServer) *server.MCPServer {
	s.AddTool(AttachTool, AttachHandler)
	return s
}
