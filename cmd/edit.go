/*
Copyright © 2025 Luke McMahon <me@lmc.id.au>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/Luke-McMahon/snip/internal/snippets"
	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit [id]",
	Short: "Edit a snippet's content or metadata",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id := args[0]
		all, err := snippets.LoadSnippets()
		if err != nil {
			return err
		}
	Use:   "edit [id]",
	Short: "Edit a snippet's content or metadata",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id := args[0]
		all, err := snippets.LoadSnippets()
		if err != nil {
			return err
		}

		var snip *snippets.Snippet
		for i := range all {
			if strings.HasPrefix(all[i].ID, id) {
				snip = &all[i]
				break
			}
		}
		var snip *snippets.Snippet
		for i := range all {
			if strings.HasPrefix(all[i].ID, id) {
				snip = &all[i]
				break
			}
		}

		if snip == nil {
			fmt.Println("Snippet not found.")
			return nil
		}
		if snip == nil {
			fmt.Println("Snippet not found.")
			return nil
		}

		tmpFile, err := os.CreateTemp("", "snip-edit-*.tmp")
		if err != nil {
			return err
		}
		defer os.Remove(tmpFile.Name())
		tmpFile, err := os.CreateTemp("", "snip-edit-*.tmp")
		if err != nil {
			return err
		}
		defer os.Remove(tmpFile.Name())

		tmpFile.WriteString(snip.Content)
		tmpFile.Close()

		editor := os.Getenv("EDITOR")
		if editor == "" {
			editor = "nano"
		}
		editor := os.Getenv("EDITOR")
		if editor == "" {
			editor = "nano"
		}

		edit := exec.Command(editor, tmpFile.Name())
		edit.Stdin = os.Stdin
		edit.Stdout = os.Stdout
		edit.Stderr = os.Stderr
		edit := exec.Command(editor, tmpFile.Name())
		edit.Stdin = os.Stdin
		edit.Stdout = os.Stdout
		edit.Stderr = os.Stderr

		if err := edit.Run(); err != nil {
			return err
		}
		if err := edit.Run(); err != nil {
			return err
		}

		b, err := os.ReadFile(tmpFile.Name())
		if err != nil {
			return err
		}
		b, err := os.ReadFile(tmpFile.Name())
		if err != nil {
			return err
		}

		snip.Content = strings.TrimRight(string(b), "\r\n")
		snip.UpdatedAt = time.Now()
		snip.Content = strings.TrimRight(string(b), "\r\n")
		snip.UpdatedAt = time.Now()

		return snippets.SaveAllSnippets(all)
	},
		return snippets.SaveAllSnippets(all)
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
	rootCmd.AddCommand(editCmd)
}
