package docker

import (
	"context"
	"fmt"
	"os/exec"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

var SBOMTool = mcp.NewTool("docker_sbom",
	mcp.WithDescription("Generates a Software Bill of Materials (SBOM) for a Docker image"),
	mcp.WithString("image",
		mcp.Required(),
		mcp.Description("The name of the image to generate SBOM for"),
	),
	mcp.WithString("format",
		mcp.Description("The format of the SBOM (e.g., spdx, cyclonedx)"),
	),
	mcp.WithString("output",
		mcp.Description("The output file for the SBOM"),
	),
)

// SBOMHandler is the handler function that handles SBOM requests
// and generates a Software Bill of Materials (SBOM) for a Docker image.
func SBOMHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	image := req.Params.Arguments["image"].(string)
	format, _ := req.Params.Arguments["format"].(string)
	output, _ := req.Params.Arguments["output"].(string)

	args := []string{"sbom", image}

	if format != "" {
		args = append(args, "--format", format)
	}
	if output != "" {
		args = append(args, "--output", output)
	}

	sbomCmd := exec.Command("docker", args...)
	outBytes, err := sbomCmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to generate SBOM for image %s: %w", image, err)
	}
	result := string(outBytes)

	return mcp.NewToolResultText(fmt.Sprintf("%s", result)), nil
}

// WithSBOMTool adds the SBOMTool to the MCP server
func WithSBOMTool(s *server.MCPServer) *server.MCPServer {
	s.AddTool(SBOMTool, SBOMHandler)
	return s
}
