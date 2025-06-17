/*
Copyright © 2025 Luke McMahon <me@lmc.id.au>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/Luke-McMahon/snip/internal/snippets"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var (
	exportFormat string
	exportOutput string
)

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export all snippets to a backup file",
	RunE: func(cmd *cobra.Command, args []string) error {
		all, err := snippets.LoadSnippets()
		if err != nil {
			return err
		}

		if len(all) == 0 {
			fmt.Println("No snippets to export")
			return nil
		}

		// Generate default filename if not provided
		if exportOutput == "" {
			timestamp := time.Now().Format("20060102-150405")
			exportOutput = fmt.Sprintf("snip-backup-%s.%s", timestamp, exportFormat)
		}

		// Ensure directory exists
		dir := filepath.Dir(exportOutput)
		if dir != "." {
			if err := os.MkdirAll(dir, 0755); err != nil {
				return fmt.Errorf("failed to create directory: %w", err)
			}
		}

		var data []byte
		switch exportFormat {
		case "json":
			data, err = json.MarshalIndent(all, "", "  ")
		case "yaml":
			data, err = yaml.Marshal(all)
		default:
			return fmt.Errorf("unsupported format: %s", exportFormat)
		}

		if err != nil {
			return fmt.Errorf("failed to marshal data: %w", err)
		}

		if err := os.WriteFile(exportOutput, data, 0644); err != nil {
			return fmt.Errorf("failed to write file: %w", err)
		}

		fmt.Printf("✅ Exported %d snippets to %s\n", len(all), exportOutput)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(exportCmd)
	exportCmd.Flags().StringVarP(&exportFormat, "format", "f", "json", "Export format (json, yaml)")
	exportCmd.Flags().StringVarP(&exportOutput, "output", "o", "", "Output file path")
}
