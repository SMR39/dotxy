package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Version is a build time variable set via -ldflags
var Version string

// versionCmd represents the version of dotxy
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of dotxy",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("dotxy version: %s\n", Version)
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
