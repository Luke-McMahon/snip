/*
Copyright © 2025 Luke McMahon <me@lmc.id.au>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/luke-mcmahon/snip/internal/gist"
	"github.com/luke-mcmahon/snip/internal/snippets"
	"github.com/spf13/cobra"
)

var gistCmd = &cobra.Command{
	Use:   "gist [id]",
	Short: "Export a snippet to GitHub Gist",
	Long: `Export a snippet to GitHub Gist.
Requires a GITHUB_TOKEN environment variable set with a valid GitHub personal access token.
The token needs 'gist' scope permissions to create gists.`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if os.Getenv("GITHUB_TOKEN") == "" {
			return fmt.Errorf("GITHUB_TOKEN environment variable not set. Please set it with a valid GitHub token")
		}

		id := args[0]
		all, err := snippets.LoadSnippets()
		if err != nil {
			return err
		}

		var snippet *snippets.Snippet
		for i := range all {
			if strings.HasPrefix(all[i].ID, id) {
				snippet = &all[i]
				break
			}
		}

		if snippet == nil {
			return fmt.Errorf("snippet with ID %s not found", id)
		}

		if err := gist.Save(snippet); err != nil {
			return fmt.Errorf("failed to save snippet to GitHub Gist: %w", err)
		}

		fmt.Println("Snippet successfully exported to GitHub Gist")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(gistCmd)
}
