/*
Copyright © 2025 Luke McMahon <me@lmc.id.au>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/luke-mcmahon/snip/internal/gist"
	"github.com/spf13/cobra"
)

var importCmd = &cobra.Command{
	Use:   "import-gists",
	Short: "Import snippets from GitHub Gists",
	Long: `Import snippets from your GitHub Gists.
Requires a GITHUB_TOKEN environment variable set with a valid GitHub personal access token.
The token needs 'gist' scope permissions to read gists.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Check if GitHub token is set
		if os.Getenv("GITHUB_TOKEN") == "" {
			return fmt.Errorf("GITHUB_TOKEN environment variable not set. Please set it with a valid GitHub token")
		}

		fmt.Println("Importing snippets from GitHub Gists...")
		if err := gist.ImportUserGists(); err != nil {
			return fmt.Errorf("failed to import gists: %w", err)
		}

		fmt.Println("Successfully imported GitHub Gists as snippets")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(importCmd)
}
