package main

import (
	"Md-converter/cmd"
	"Md-converter/cmd/convert"
	"Md-converter/logger"
)

func main() {

	rootCmd := cmd.RootCmd()

	convert.Init(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		logger.HaltOnError(err, "Initial setup failed")
	}

}
