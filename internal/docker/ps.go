package docker

import (
	"context"
	"fmt"
	"os/exec"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// PSTool is the tool definition for the Docker ps command.
var PSTool = mcp.NewTool("docker_ps",
	mcp.WithDescription("Lists all running docker containers"),
	mcp.WithString("filter",
		mcp.Description("Filter output based on conditions provided"),
	),
	mcp.WithString("all",
		mcp.Description("Show all containers (default shows just running)"),
	),
	mcp.WithString("format",
		mcp.Description("Format the output using a custom template"),
	),
	mcp.WithString("latest",
		mcp.Description("Show the latest created container (includes all states)"),
	),
	mcp.WithString("no-trunc",
		mcp.Description("Don't truncate output"),
	),
)

// PSHandler is the handler function that handles ps requests, to list
// all running containers, while using docker cli tool.
func PSHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := []string{"ps"}
	filter, exist := req.Params.Arguments["filter"]
	if exist {
		args = append(args, "--filter", filter.(string))
	}

	_, exist = req.Params.Arguments["all"]
	if exist {
		args = append(args, "--all")
	}

	format, exist := req.Params.Arguments["format"]
	if exist {
		args = append(args, "--format", format.(string))
	}

	_, exist = req.Params.Arguments["latest"]
	if exist {
		args = append(args, "--latest")
	}

	_, exist = req.Params.Arguments["no-trunc"]
	if exist {
		args = append(args, "--no-trunc")
	}

	psCmd := exec.Command("docker", args...)
	outBytes, err := psCmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to list containers: %w", err)
	}
	result := string(outBytes)

	return mcp.NewToolResultText(fmt.Sprintf("%s", result)), nil
}

// WithPSTool is a convenience function to add the Docker ps tool to the MCP server.
func WithPSTool(s *server.MCPServer) *server.MCPServer {
	s.AddTool(PSTool, PSHandler)
	return s
}
