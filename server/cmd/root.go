package cmd

import (
	"os"

	publisherServer "github.com/amaany3/minivr/server/cmd/publisher-server"
	"github.com/amaany3/minivr/server/internal/logger"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "minivr"}

func init() {
	rootCmd.AddCommand(publisherServer.Cmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logger.Error("failed to execute root command: %v", err)
		os.Exit(1)
	}
}
