package docker

import (
	"context"
	"fmt"
	"os/exec"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// DiffTool is the tool definition for the Docker diff command.
var DiffTool = mcp.NewTool("docker_diff",
	mcp.WithDescription("Shows the changes made to a container's filesystem"),
	mcp.WithString("containerID",
		mcp.Required(),
		mcp.Description("The ID of the container to show changes for"),
	),
)

// DiffHandler is the handler function that handles diff requests
// and actually makes a use of docker cli tool to show the changes made
func DiffHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	containerID := req.Params.Arguments["containerID"].(string)

	diffCmd := exec.Command("docker", "diff", containerID)
	outBytes, err := diffCmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to show changes for container %s: %w", containerID, err)
	}
	result := string(outBytes)

	return mcp.NewToolResultText(fmt.Sprintf("%s", result)), nil
}

// WithDiffTool adds the DiffTool to the MCP server
func WithDiffTool(s *server.MCPServer) *server.MCPServer {
	s.AddTool(DiffTool, DiffHandler)
	return s
}
