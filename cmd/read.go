/*
Copyright © 2025 Luke McMahon <me@lmc.id.au>
*/
package cmd

import (
	"fmt"

	"github.com/Luke-McMahon/snip/internal/snippets"
	"github.com/spf13/cobra"
)

var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Read a snipper by ID",
	RunE: func(cmd *cobra.Command, args []string) error {
		id := args[0]

		snip, err := snippets.FindSnippetByID(id)
		if err != nil {
			return fmt.Errorf("snippet not found: %s", id)
		}

		fmt.Printf("📎  %s\n", snip.Title)
		if len(snip.Tags) > 0 {
			fmt.Printf("🏷️  Tags: %v\n", snip.Tags)
		}
		if snip.Language != "" {
			fmt.Printf("🧠  Language: %s\n", snip.Language)
		}
		fmt.Printf("⭐  Starred: %v\n🔐  Private: %v\n", snip.Starred, snip.Private)
		fmt.Printf("📅  Created: %s\n", snip.CreatedAt.Format("2006-01-02 15:04"))
		fmt.Printf("📎  %s\n", snip.Title)
		if len(snip.Tags) > 0 {
			fmt.Printf("🏷️  Tags: %v\n", snip.Tags)
		}
		if snip.Language != "" {
			fmt.Printf("🧠  Language: %s\n", snip.Language)
		}
		fmt.Printf("⭐  Starred: %v\n🔐  Private: %v\n", snip.Starred, snip.Private)
		fmt.Printf("📅  Created: %s\n", snip.CreatedAt.Format("2006-01-02 15:04"))

		fmt.Println("\n📝 Snippet:\n")
		fmt.Println(snip.Content)
		fmt.Println("\n📝 Snippet:\n")
		fmt.Println(snip.Content)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(readCmd)
}
