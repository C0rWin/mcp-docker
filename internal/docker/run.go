package docker

import (
	"context"
	"fmt"
	"os/exec"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

var RunTool = mcp.NewTool("docker_run",
	mcp.WithDescription("Runs a command in a new container"),
	mcp.WithString("image",
		mcp.Required(),
		mcp.Description("The name of the image to run"),
	),
	mcp.WithString("command",
		mcp.Description("The command to run in the container"),
	),
	mcp.WithString("name",
		mcp.Description("The name to assign to the container"),
	),
	mcp.WithString("interactive",
		mcp.Description("Run the container in interactive mode"),
	),
	mcp.WithString("rm",
		mcp.Description("Automatically remove the container when it exits"),
	),
	mcp.WithString("detach",
		mcp.Description("Run the container in detached mode"),
	),
	mcp.WithString("workdir",
		mcp.Description("Set the working directory inside the container"),
	),
	mcp.WithString("network",
		mcp.Description("Connect the container to a network"),
	),
	mcp.WithString("env",
		mcp.Description("Set environment variables in the container"),
	),
	mcp.WithString("volume",
		mcp.Description("Mount a volume into the container"),
	),
)

// RunHandler is the handler function that handles run requests
func RunHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	image := req.Params.Arguments["image"].(string)
	command, _ := req.Params.Arguments["command"].(string)
	name, _ := req.Params.Arguments["name"].(string)
	interactive, _ := req.Params.Arguments["interactive"].(string)
	rm, _ := req.Params.Arguments["rm"].(string)
	detach, _ := req.Params.Arguments["detach"].(string)
	workdir, _ := req.Params.Arguments["workdir"].(string)
	network, _ := req.Params.Arguments["network"].(string)
	env, _ := req.Params.Arguments["env"].(string)
	volume, _ := req.Params.Arguments["volume"].(string)

	args := []string{"run"}

	if name != "" {
		args = append(args, "--name", name)
	}
	if interactive != "" {
		args = append(args, "-it")
	}
	if rm != "" {
		args = append(args, "--rm")
	}
	if detach != "" {
		args = append(args, "-d")
	}
	if workdir != "" {
		args = append(args, "--workdir", workdir)
	}
	if network != "" {
		args = append(args, "--network", network)
	}
	if env != "" {
		args = append(args, "--env", env)
	}
	if volume != "" {
		args = append(args, "--volume", volume)
	}

	args = append(args, image)

	if command != "" {
		// Split the command string into individual arguments
		cmdArgs := strings.Fields(command)
		args = append(args, cmdArgs...)
	}

	runCmd := exec.Command("docker", args...)
	outBytes, err := runCmd.Output()
	if err != nil {
		return nil, fmt.Errorf(`failed to run container, "[%s]": %w`,
			strings.Join(runCmd.Args, " "), err)
	}
	result := string(outBytes)

	return mcp.NewToolResultText(fmt.Sprintf("%s", result)), nil
}

// WithRunTool adds the RunTool to the MCP server
func WithRunTool(s *server.MCPServer) *server.MCPServer {
	s.AddTool(RunTool, RunHandler)
	return s
}
