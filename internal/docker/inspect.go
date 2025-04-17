package docker

import (
	"context"
	"fmt"
	"os/exec"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// In this file, we define the inspect tool for Docker,
// which allows users to inspect a container, image, or volume,
// while using docker cli tool.
var InspectTool = mcp.NewTool("docker_inspect",
	mcp.WithDescription("Inspects a docker container, image, or volume"),
	mcp.WithString("containerID",
		mcp.Required(),
		mcp.Description("The ID of the container to inspect"),
	),
)

// InspectHandler is the handler function that handles inspection requests
// and actually makes a use of docker cli tool to inspect the container.
func InspectHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	containerID := req.Params.Arguments["containerID"].(string)

	inspectCmd := exec.Command("docker", "inspect", containerID)
	outBytes, err := inspectCmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to inspect container: %w", err)
	}
	result := string(outBytes)

	return mcp.NewToolResultText(fmt.Sprintf("%s", result)), nil
}

func WithInspectTool(s *server.MCPServer) *server.MCPServer {
	s.AddTool(InspectTool, InspectHandler)
	return s
}
