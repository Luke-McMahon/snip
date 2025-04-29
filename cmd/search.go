/*
Copyright © 2025 Luke McMahon <me@lmc.id.au>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/luke-mcmahon/snip/internal/snippets"
	"github.com/spf13/cobra"
)

var (
	tagFilter string
	languageFilter string
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search [query]",
	Short: "Search snippets by title, content, or tags",
	RunE: func(cmd *cobra.Command, args []string) error {
		query := strings.ToLower(strings.Join(args, " "))
		all, err := snippets.LoadSnippets()
		if err != nil {
			return err
		}

		found := 0
		for _, snip := range all {
			title := strings.ToLower(snip.Title)
			content := strings.ToLower(snip.Content)
			tags := strings.ToLower(strings.Join(snip.Tags, ","))

			if strings.Contains(title, query) || strings.Contains(content, query) || strings.Contains(tags, query) {
				if tagFilter != "" && !containsIgnoreCase(snip.Tags, tagFilter) {
					continue
				}

				if languageFilter != "" && !strings.EqualFold(snip.Language, languageFilter) {
					continue
				}
				fmt.Printf("• %s [%s]\n\tID: %s\n\n", snip.Title, strings.Join(snip.Tags, ", "), snip.ID)
				found++
			}
		}

		if found == 0 {
			fmt.Println("No matching snippets found.")
		}

		return nil
	},
}

func containsIgnoreCase(slice []string, match string) bool {
	match = strings.ToLower(match)
	for _, s := range slice {
		if strings.ToLower(s) == match {
			return true
		}
	}
	return false
}

func init() {
	rootCmd.AddCommand(searchCmd)
    searchCmd.Flags().StringVar(&tagFilter, "tag", "", "Filter by tag")
    searchCmd.Flags().StringVar(&languageFilter, "lang", "", "Filter by language")
}
