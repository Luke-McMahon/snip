/*
Copyright © 2025 Luke McMahon <me@lmc.id.au>
*/
package cmd

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"

	"github.com/Luke-McMahon/snip/internal/snippets"
	"github.com/spf13/cobra"
)

var copyCmd = &cobra.Command{
	Use:   "copy [id]",
	Short: "Copy a snippet's content to clipboard",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id := args[0]

		snip, err := snippets.FindSnippetByID(id)
		if err != nil {
			return fmt.Errorf("snippet not found: %s", id)
		}

		// Copy to clipboard based on OS
		var copyCmd *exec.Cmd
		switch runtime.GOOS {
		case "darwin":
			copyCmd = exec.Command("pbcopy")
		case "linux":
			// Try xclip first, then xsel
			if _, err := exec.LookPath("xclip"); err == nil {
				copyCmd = exec.Command("xclip", "-selection", "clipboard")
			} else if _, err := exec.LookPath("xsel"); err == nil {
				copyCmd = exec.Command("xsel", "--clipboard", "--input")
			} else {
				return fmt.Errorf("clipboard access requires xclip or xsel on Linux")
			}
		case "windows":
			copyCmd = exec.Command("clip")
		default:
			return fmt.Errorf("clipboard copy not supported on %s", runtime.GOOS)
		}

		// Write content to clipboard
		copyCmd.Stdin = strings.NewReader(snip.Content)
		if err := copyCmd.Run(); err != nil {
			return fmt.Errorf("failed to copy to clipboard: %w", err)
		}

		fmt.Printf("📋 Copied snippet '%s' to clipboard\n", snip.Title)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(copyCmd)
}
