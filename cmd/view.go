/*
Copyright © 2025 Luke McMahon <me@lmc.id.au>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/Luke-McMahon/snip/internal/snippets"
	"github.com/alecthomas/chroma/quick"
	"github.com/spf13/cobra"
)

var disableSyntax bool

var viewCmd = &cobra.Command{
	Use:   "view [id]",
	Short: "View a full snippet by its ID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id := args[0]
		all, err := snippets.LoadSnippets()
		if err != nil {
			return err
		}

		for _, snip := range all {
			if strings.HasPrefix(snip.ID, id) {
				printSnippet(snip)
				return nil
			}
		}

		fmt.Println("Snippet not found.")
		return nil
	},
}

func printSnippet(s snippets.Snippet) {
	fmt.Println("────────────────────────────")
	fmt.Printf("📌 %s [%s]\n", s.Title, strings.Join(s.Tags, ", "))
	fmt.Printf("🆔 %s\n", s.ID)
	if s.Language != "" {
		fmt.Printf("🗣️  Language: %s\n", s.Language)
	}
	if s.Starred {
		fmt.Println("⭐ Starred")
	}
	if s.Private {
		fmt.Println("🔒 Private")
	}
	fmt.Printf("🕒 Created: %s\n", s.CreatedAt.Format("2006-01-02 15:04"))
	fmt.Printf("🕓 Updated: %s\n", s.UpdatedAt.Format("2006-01-02 15:04"))
	fmt.Println("────────────────────────────")
	renderContent(s)
	fmt.Println("────────────────────────────")
}

func renderContent(s snippets.Snippet) {
	if disableSyntax {
		fmt.Println(s.Content)
		return
	}

	err := quick.Highlight(os.Stdout, s.Content, s.Language, "terminal16m", "monokai")
	if err != nil {
		// Fallback if highlighting fails
		fmt.Println(s.Content)
	} else {
		fmt.Println()
	}
}

func init() {
	rootCmd.AddCommand(viewCmd)
	viewCmd.Flags().BoolVar(&disableSyntax, "no-highlight", false, "Display without syntax highlighting")
}
