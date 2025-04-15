/*
Copyright © 2025 Luke McMahon <me@lmc.id.au>

*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
    "github.com/luke-mcmahon/snipit/internal/snippets"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all saved snippets",
	RunE: func(cmd *cobra.Command, args []string) error {
	  all, err := snippets.LoadSnippets()
	  if err != nil {
		  return err
	  }

	if len(all) == 0 {
		fmt.Println("No snippets saved yet.")
		return nil
	}

	for _, snip := range all {
	  tagStr := ""
	  if len(snip.Tags) > 0 {
		tagStr = "[" + strings.Join(snip.Tags, ", ") + "]"
	  }

	  fmt.Printf("- %s %s\n\tID: %s\n\n", snip.Title, tagStr, snip.ID)
	}
	return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
