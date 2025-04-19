package docker

import (
	"context"
	"fmt"
	"os/exec"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

var CommitTool = mcp.NewTool("docker_commit",
	mcp.WithDescription("Creates a new image from a container's changes"),
	mcp.WithString("containerID",
		mcp.Required(),
		mcp.Description("The ID of the container to commit"),
	),
	mcp.WithString("repository",
		mcp.Description("The repository name for the new image"),
	),
	mcp.WithString("tag",
		mcp.Description("The tag for the new image"),
	),
	mcp.WithString("message",
		mcp.Description("A commit message"),
	),
	mcp.WithString("author",
		mcp.Description("The author of the new image"),
	),
	mcp.WithString("change",
		mcp.Description("Apply a Dockerfile instruction to the container's filesystem"),
	),
	mcp.WithString("pause",
		mcp.Description("Pause the container during commit"),
	),
)

// CommitHandler is the handler function that handles commit requests
// and actually makes a use of docker cli tool to commit the container.
func CommitHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	containerID := req.Params.Arguments["containerID"].(string)

	args := []string{"commit", containerID}

	repository, exist := req.Params.Arguments["repository"]
	if exist {
		args = append(args, repository.(string))
	}

	tag, exist := req.Params.Arguments["tag"]
	if exist {
		args = append(args, tag.(string))
	}

	message, exist := req.Params.Arguments["message"]
	if exist {
		args = append(args, "--message", message.(string))
	}

	author, exist := req.Params.Arguments["author"]
	if exist {
		args = append(args, "--author", author.(string))
	}

	change, exist := req.Params.Arguments["change"]
	if exist {
		args = append(args, "--change", change.(string))
	}

	pause, exist := req.Params.Arguments["pause"]
	if exist {
		args = append(args, "--pause", pause.(string))
	}

	commitCmd := exec.Command("docker", args...)
	outBytes, err := commitCmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to commit container: %w", err)
	}
	result := string(outBytes)

	return mcp.NewToolResultText(fmt.Sprintf("%s", result)), nil
}

// WithCommitTool adds the commit tool to the MCP server
func WithCommitTool(s *server.MCPServer) *server.MCPServer {
	s.AddTool(CommitTool, CommitHandler)
	return s
}
