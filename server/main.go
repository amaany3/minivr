package main

import (
	"github.com/amaany3/minivr/server/cmd"
	"github.com/amaany3/minivr/server/internal/logger"
)

func init() {
	logger.EnableCloudLoggingLogger()
}

func main() {
	cmd.Execute()
}
