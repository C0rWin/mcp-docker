/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/mark3labs/mcp-docker/internal/docker"
	"github.com/mark3labs/mcp-go/server"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		s := server.NewMCPServer(
			"Calculator Demo Application",
			"1.0.0",
			server.WithResourceCapabilities(true, true),
			server.WithPromptCapabilities(true),
			server.WithLogging(),
			server.WithRecovery(),
		)

		docker.WithInspectTool(s)
		docker.WithPSTool(s)
		docker.WithHistoryTool(s)
		docker.WithDiffTool(s)
		docker.WithRunTool(s)
		docker.WithExecTool(s)
		docker.WithSBOMTool(s)
		docker.WithImageTools(s)
		docker.WithSearchTool(s)
		docker.WithPullTool(s)

		if err := server.ServeStdio(s); err != nil {
			fmt.Println("Error starting server:", err)
		}

	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
