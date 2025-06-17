/*
Copyright © 2025 Luke McMahon <me@lmc.id.au>
*/
package cmd

import (
	"fmt"

	"github.com/Luke-McMahon/snip/internal/snippets"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete a snippet by ID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id := args[0]

		// Find the snippet first to confirm it exists
		snip, err := snippets.FindSnippetByID(id)
		if err != nil {
			return fmt.Errorf("snippet not found: %s", id)
		}

		// Load all snippets
		all, err := snippets.LoadSnippets()
		if err != nil {
			return err
		}

		// Remove the snippet from the slice
		var filtered []snippets.Snippet
		for _, s := range all {
			if s.ID != snip.ID {
				filtered = append(filtered, s)
			}
		}

		// Save the updated slice
		if err := snippets.SaveAllSnippets(filtered); err != nil {
			return err
		}

		fmt.Printf("Deleted snippet: %s\n", snip.Title)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
