/*
Copyright © 2025 Luke McMahon <me@lmc.id.au>
*/
package cmd

import (
	"fmt"

	"github.com/luke-mcmahon/snip/internal/snippets"
	"github.com/spf13/cobra"
)

// readCmd represents the read command
var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Read a snipper by ID",
	RunE: func(cmd *cobra.Command, args []string) error {
		id := args[0]

	snip, err := snippets.FindSnippetByID(id)
	if err != nil { 
fmt.Errorf("snippet not found: %s", id)
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

        fmt.Println("\n📝 Snippet:\n")
        fmt.Println(snip.Content)

        return nil
	},
}

func init() {
	rootCmd.AddCommand(readCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// readCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// readCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
