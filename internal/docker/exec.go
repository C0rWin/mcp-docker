package docker

import (
	"context"
	"fmt"
	"os/exec"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

var ExecTool = mcp.NewTool("docker_exec",
	mcp.WithDescription("Executes a command in a running container"),
	mcp.WithString("containerID",
		mcp.Required(),
		mcp.Description("The ID of the container to execute the command in"),
	),
	mcp.WithString("command",
		mcp.Required(),
		mcp.Description("The command to execute in the container"),
	),
	mcp.WithString("interactive",
		mcp.Description("Run the command in interactive mode"),
	),
	mcp.WithString("detach",
		mcp.Description("Run the command in detached mode"),
	),
)

// ExecHandler is the handler function that handles exec requests
func ExecHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	containerID := req.Params.Arguments["containerID"].(string)
	command := req.Params.Arguments["command"].(string)
	interactive, _ := req.Params.Arguments["interactive"].(string)
	detach, _ := req.Params.Arguments["detach"].(string)

	args := []string{"exec", containerID}
	if interactive != "" {
		args = append(args, "-i")
	}
	if detach != "" {
		args = append(args, "-d")
	}
	// Split the command string into individual arguments
	cmdArgs := strings.Fields(command)
	args = append(args, cmdArgs...)

	execCmd := exec.Command("docker", args...)
	outBytes, err := execCmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to execute command in container %s, \"[%s]\": %w",
			containerID, strings.Join(execCmd.Args, " "), err)
	}
	result := string(outBytes)

	return mcp.NewToolResultText(result), nil
}

func WithExecTool(s *server.MCPServer) *server.MCPServer {
	s.AddTool(ExecTool, ExecHandler)
	return s
}
