package main

import "github.com/amaany3/minivr/server/internal/logger"

func init() {
	logger.EnableCloudLoggingLogger()
}

func main() {
	logger.Info("hello world!")
}
