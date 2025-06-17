/*
Copyright © 2025 Luke McMahon <me@lmc.id.au>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/Luke-McMahon/snip/internal/snippets"
	"github.com/spf13/cobra"
)

var starCmd = &cobra.Command{
	Use:   "star [id]",
	Short: "Star a snippet",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id := args[0]

		// Find the snippet
		snip, err := snippets.FindSnippetByID(id)
		if err != nil {
			return fmt.Errorf("snippet not found: %s", id)
		}

		if snip.Starred {
			fmt.Printf("Snippet '%s' is already starred\n", snip.Title)
			return nil
		}

		// Load all snippets
		all, err := snippets.LoadSnippets()
		if err != nil {
			return err
		}

		// Update the specific snippet
		for i := range all {
			if all[i].ID == snip.ID {
				all[i].Starred = true
				all[i].UpdatedAt = time.Now()
				break
			}
		}

		// Save the updated slice
		if err := snippets.SaveAllSnippets(all); err != nil {
			return err
		}

		fmt.Printf("⭐ Starred snippet: %s\n", snip.Title)
		return nil
	},
}

var unstarCmd = &cobra.Command{
	Use:   "unstar [id]",
	Short: "Unstar a snippet",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id := args[0]

		// Find the snippet
		snip, err := snippets.FindSnippetByID(id)
		if err != nil {
			return fmt.Errorf("snippet not found: %s", id)
		}

		if !snip.Starred {
			fmt.Printf("Snippet '%s' is not starred\n", snip.Title)
			return nil
		}

		// Load all snippets
		all, err := snippets.LoadSnippets()
		if err != nil {
			return err
		}

		// Update the specific snippet
		for i := range all {
			if all[i].ID == snip.ID {
				all[i].Starred = false
				all[i].UpdatedAt = time.Now()
				break
			}
		}

		// Save the updated slice
		if err := snippets.SaveAllSnippets(all); err != nil {
			return err
		}

		fmt.Printf("⭐ Unstarred snippet: %s\n", snip.Title)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(starCmd)
	rootCmd.AddCommand(unstarCmd)
}
