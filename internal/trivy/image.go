package trivy

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

var ImageTool = mcp.NewTool("trivy_image",
	mcp.WithDescription("Scan a Docker image for vulnerabilities"),
	mcp.WithString("image",
		mcp.Required(),
		mcp.Description("The name of the image to scan"),
	),
)

// ImageHandler is the handler function that handles image scan requests
// and actually makes a use of trivy cli tool to scan the image.
func ImageHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	image := req.Params.Arguments["image"].(string)

	// Run the trivy command to scan the image
	var stderr bytes.Buffer
	scanCmd := exec.Command("trivy", "image", image)
	scanCmd.Stderr = &stderr
	outBytes, err := scanCmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to scan image: %w \n %s", err, stderr.String())
	}
	result := string(outBytes)

	return mcp.NewToolResultText(fmt.Sprintf("%s", result)), nil
}

// WithImageTool adds the image tool to the given mcp server
func WithImageTool(srv *server.MCPServer) *server.MCPServer {
	srv.AddTool(ImageTool, ImageHandler)
	return srv
}
