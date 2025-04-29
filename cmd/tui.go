/*
Copyright © 2025 Luke McMahon <me@lmc.id.au>

*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/luke-mcmahon/snip/internal/tui"
)

var tuiCmd = &cobra.Command{
  Use:   "tui",
  Short: "Launch the interactive TUI snippet browser",
  RunE: func(cmd *cobra.Command, args []string) error {
    return tui.Run()  // calls tea.NewProgram(...).Start()
  },
}

func init() {
  rootCmd.AddCommand(tuiCmd)
}
