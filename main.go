package main

import (
	"github.com/Zizu-oswald/MdConverter/cmd"
	"github.com/Zizu-oswald/MdConverter/cmd/convert"
	"github.com/Zizu-oswald/MdConverter/logger"
)

func main() {

	rootCmd := cmd.RootCmd()

	convert.Init(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		logger.HaltOnError(err, "Initial setup failed")
	}

}
