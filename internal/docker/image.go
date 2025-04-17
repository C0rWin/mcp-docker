package docker

import (
	"context"
	"fmt"
	"os/exec"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

var ImageListTools = mcp.NewTool("docker_image",
	mcp.WithDescription("Docker image management commands"),
	mcp.WithString("list",
		mcp.Description("List all images"),
	),
)

var ImageInspectTool = mcp.NewTool("docker_image_inspect",
	mcp.WithDescription("Inspect a Docker image"),
	mcp.WithString("imageID",
		mcp.Required(),
		mcp.Description("The ID of the image to inspect"),
	),
	mcp.WithString("format",
		mcp.Description("Format the output using a Go template"),
	),
	mcp.WithString("size",
		mcp.Description("Display image size"),
	),
)

var ImageHistoryTool = mcp.NewTool("docker_image_history",
	mcp.WithDescription("Show the history of an image"),
	mcp.WithString("imageID",
		mcp.Required(),
		mcp.Description("The ID of the image to show history for"),
	),
	mcp.WithString("format",
		mcp.Description("Format the output using a Go template"),
	),
	mcp.WithString("no-trunc",
		mcp.Description("Don't truncate output"),
	),
)

// ImageListHandler is the handler function that handles image requests
func ImageListHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// Implement the logic to list Docker images
	diffCmd := exec.Command("docker", "image", "ls")
	outBytes, err := diffCmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to list images: %w", err)
	}
	result := string(outBytes)

	return mcp.NewToolResultText(fmt.Sprintf("%s", result)), nil
}

// ImageInspectHandler is the handler function that handles image inspect requests
func ImageInspectHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// Implement the logic to inspect a Docker image
	// This is a placeholder implementation
	imageID := req.Params.Arguments["imageID"].(string)
	args := []string{"image", "inspect", imageID}
	_, exist := req.Params.Arguments["size"]
	if exist {
		args = append(args, "--size")
	}

	format, exist := req.Params.Arguments["format"]
	if exist {
		args = append(args, "--format", format.(string))
	}
	diffCmd := exec.Command("docker", args...)
	outBytes, err := diffCmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to inspect image %s: %w", imageID, err)
	}
	result := string(outBytes)
	if result == "" {
		return nil, fmt.Errorf("no image found with ID %s", imageID)
	}
	return mcp.NewToolResultText(fmt.Sprintf("%s", result)), nil
}

// ImageHistoryHandler is the handler function that handles image history requests
func ImageHistoryHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// Implement the logic to show the history of a Docker image
	// This is a placeholder implementation
	imageID := req.Params.Arguments["imageID"].(string)
	args := []string{"image", "history", imageID}
	format, exist := req.Params.Arguments["format"]
	if exist {
		args = append(args, "--format", format.(string))
	}
	_, exist = req.Params.Arguments["no-trunc"]
	if exist {
		args = append(args, "--no-trunc")
	}

	diffCmd := exec.Command("docker", args...)
	outBytes, err := diffCmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to show history for image %s: %w", imageID, err)
	}
	result := string(outBytes)
	if result == "" {
		return nil, fmt.Errorf("no history found for image %s", imageID)
	}
	return mcp.NewToolResultText(fmt.Sprintf("%s", result)), nil
}

func WithImageTools(s *server.MCPServer) *server.MCPServer {
	s.AddTool(ImageListTools, ImageListHandler)
	s.AddTool(ImageInspectTool, ImageInspectHandler)
	s.AddTool(ImageHistoryTool, ImageHistoryHandler)
	return s
}
