/*
Copyright © 2025 Luke McMahon <me@lmc.id.au>
*/
package cmd

import (
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/Luke-McMahon/snip/internal/gist"
	"github.com/Luke-McMahon/snip/internal/snippets"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var (
	content  string
	tags     string
	language string
	starred  bool
	private  bool
	remote   bool
)

var addCmd = &cobra.Command{
	Use:   "add [title]",
	Short: "Add a new snippet, optionally pushing to a remote service",
	Args:  cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		title := args[0]
		remote, err := cmd.Flags().GetBool("remote")
		if err != nil {
			return err
		}
		snippetContent := content

		if snippetContent == "" {
			tmpFile, err := os.CreateTemp("", "snip-*.tmp")
			if err != nil {
				return err
			}
			defer os.Remove(tmpFile.Name())

			editor := os.Getenv("EDITOR")
			if editor == "" {
				editor = "nano"
			}

			edit := exec.Command(editor, tmpFile.Name())
			edit.Stdin = os.Stdin
			edit.Stdout = os.Stdout
			edit.Stderr = os.Stderr

			if err := edit.Run(); err != nil {
				return err
			}

			b, err := os.ReadFile(tmpFile.Name())
			if err != nil {
				return err
			}

			snippetContent = string(b)
			snippetContent = strings.TrimRight(snippetContent, "\r\n")
		}

		snippet := snippets.Snippet{
			ID:        uuid.NewString(),
			Title:     title,
			Tags:      strings.Split(tags, ","),
			Content:   snippetContent,
			Language:  language,
			Starred:   starred,
			Private:   private,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		if remote {
			return gist.Save(&snippet)
		}

		return snippets.SaveSnippet(snippet)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVar(&content, "content", "", "Snippet content")
	addCmd.Flags().StringVar(&tags, "tags", "", "Comma-separated tags")
	addCmd.Flags().StringVar(&language, "language", "", "Snippet language")
	addCmd.Flags().BoolVar(&starred, "starred", false, "Mark snippet as starred")
	addCmd.Flags().BoolVar(&private, "private", false, "Mark snippet as private")
	addCmd.Flags().BoolVar(&remote, "remote", false, "Push snippet to remote service")
}
