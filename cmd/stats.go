/*
Copyright © 2025 Luke McMahon <me@lmc.id.au>
*/
package cmd

import (
	"fmt"
	"sort"

	"github.com/Luke-McMahon/snip/internal/snippets"
	"github.com/spf13/cobra"
)

type TagCount struct {
	Tag   string
	Count int
}

type LanguageCount struct {
	Language string
	Count    int
}

var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Show statistics about your snippets",
	RunE: func(cmd *cobra.Command, args []string) error {
		all, err := snippets.LoadSnippets()
		if err != nil {
			return err
		}

		if len(all) == 0 {
			fmt.Println("No snippets found")
			return nil
		}

		// Count tags
		tagCounts := make(map[string]int)
		languageCounts := make(map[string]int)
		starredCount := 0
		privateCount := 0

		for _, snip := range all {
			// Count tags
			for _, tag := range snip.Tags {
				tagCounts[tag]++
			}

			// Count languages
			if snip.Language != "" {
				languageCounts[snip.Language]++
			}

			// Count starred/private
			if snip.Starred {
				starredCount++
			}
			if snip.Private {
				privateCount++
			}
		}

		// Sort tags by count
		var tags []TagCount
		for tag, count := range tagCounts {
			tags = append(tags, TagCount{Tag: tag, Count: count})
		}
		sort.Slice(tags, func(i, j int) bool {
			return tags[i].Count > tags[j].Count
		})

		// Sort languages by count
		var languages []LanguageCount
		for lang, count := range languageCounts {
			languages = append(languages, LanguageCount{Language: lang, Count: count})
		}
		sort.Slice(languages, func(i, j int) bool {
			return languages[i].Count > languages[j].Count
		})

		// Display stats
		fmt.Printf("📊 Snippet Statistics\n")
		fmt.Printf("─────────────────────\n")
		fmt.Printf("📎 Total snippets: %d\n", len(all))
		fmt.Printf("⭐ Starred: %d\n", starredCount)
		fmt.Printf("🔐 Private: %d\n", privateCount)

		if len(tags) > 0 {
			fmt.Printf("\n🏷️  Top Tags:\n")
			limit := 10
			if len(tags) < limit {
				limit = len(tags)
			}
			for i := 0; i < limit; i++ {
				fmt.Printf("   %s: %d\n", tags[i].Tag, tags[i].Count)
			}
		}

		if len(languages) > 0 {
			fmt.Printf("\n🧠 Top Languages:\n")
			limit := 10
			if len(languages) < limit {
				limit = len(languages)
			}
			for i := 0; i < limit; i++ {
				fmt.Printf("   %s: %d\n", languages[i].Language, languages[i].Count)
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(statsCmd)
}
