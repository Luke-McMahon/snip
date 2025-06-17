/*
Copyright © 2025 Luke McMahon <me@lmc.id.au>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Luke-McMahon/snip/internal/snippets"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var (
	importMerge bool
)

var importCmd = &cobra.Command{
	Use:   "import [file]",
	Short: "Import snippets from a backup file",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		filename := args[0]

		// Check if file exists
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			return fmt.Errorf("file not found: %s", filename)
		}

		// Read file
		data, err := os.ReadFile(filename)
		if err != nil {
			return fmt.Errorf("failed to read file: %w", err)
		}

		// Determine format from extension
		ext := strings.ToLower(filepath.Ext(filename))
		var importedSnippets []snippets.Snippet

		switch ext {
		case ".json":
			err = json.Unmarshal(data, &importedSnippets)
		case ".yaml", ".yml":
			err = yaml.Unmarshal(data, &importedSnippets)
		default:
			return fmt.Errorf("unsupported file format: %s (supported: .json, .yaml, .yml)", ext)
		}

		if err != nil {
			return fmt.Errorf("failed to parse file: %w", err)
		}

		if len(importedSnippets) == 0 {
			fmt.Println("No snippets found in file")
			return nil
		}

		var allSnippets []snippets.Snippet
		if importMerge {
			// Load existing snippets
			existing, err := snippets.LoadSnippets()
			if err != nil {
				return err
			}
			allSnippets = existing
		}

		// Check for duplicates and merge
		existingIDs := make(map[string]bool)
		for _, snip := range allSnippets {
			existingIDs[snip.ID] = true
		}

		imported := 0
		skipped := 0
		for _, snip := range importedSnippets {
			if existingIDs[snip.ID] {
				skipped++
				continue
			}
			allSnippets = append(allSnippets, snip)
			imported++
		}

		// Save all snippets
		if err := snippets.SaveAllSnippets(allSnippets); err != nil {
			return err
		}

		fmt.Printf("✅ Imported %d snippets", imported)
		if skipped > 0 {
			fmt.Printf(" (%d skipped - already exist)", skipped)
		}
		fmt.Printf(" from %s\n", filename)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(importCmd)
	importCmd.Flags().BoolVarP(&importMerge, "merge", "m", true, "Merge with existing snippets (default: true)")
}
