package cmd

import (
	"github.com/spf13/cobra"
)

func RootCmd() *cobra.Command {
	return &cobra.Command{
		Use: "mdtool",
		Short: "Tool for MarkDown files",
		Long: "Tool for MarkDown files",
	}
}

