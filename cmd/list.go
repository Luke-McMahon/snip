/*
Copyright © 2025 Luke McMahon <me@lmc.id.au>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/Luke-McMahon/snip/internal/snippets"
	"github.com/spf13/cobra"
)

var (
	listStarred bool
	listPrivate bool
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all saved snippets",
	RunE: func(cmd *cobra.Command, args []string) error {
		all, err := snippets.LoadSnippets()
		if err != nil {
			return err
		}

		// Filter snippets based on flags
		var filtered []snippets.Snippet
		for _, snip := range all {
			if listStarred && !snip.Starred {
				continue
			}
			if listPrivate && !snip.Private {
				continue
			}
			filtered = append(filtered, snip)
		}

		if len(filtered) == 0 {
			if listStarred {
				fmt.Println("No starred snippets found.")
			} else if listPrivate {
				fmt.Println("No private snippets found.")
			} else {
				fmt.Println("No snippets saved yet.")
			}
			return nil
		}

		for _, snip := range filtered {
			tagStr := ""
			if len(snip.Tags) > 0 {
				tagStr = "[" + strings.Join(snip.Tags, ", ") + "]"
			}

			starIcon := ""
			if snip.Starred {
				starIcon = "⭐ "
			}
			privateIcon := ""
			if snip.Private {
				privateIcon = "🔐 "
			}

			fmt.Printf("- %s%s%s %s\n\tID: %s\n\n", starIcon, privateIcon, snip.Title, tagStr, snip.ID)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVar(&listStarred, "starred", false, "Show only starred snippets")
	listCmd.Flags().BoolVar(&listPrivate, "private", false, "Show only private snippets")
}
