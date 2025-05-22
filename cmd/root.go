/*
Copyright © 2025 Luke McMahon <me@lmc.id.au>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "snip",
	Short: "snip is a CLI snippet manager",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
