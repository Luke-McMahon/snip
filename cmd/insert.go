/*
Copyright © 2025 Luke McMahon <me@lmc.id.au>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Luke-McMahon/snip/internal/snippets"
	"github.com/Luke-McMahon/snip/internal/templates"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var (
	insertList bool
)

var insertCmd = &cobra.Command{
	Use:   "insert [template-name] [title]",
	Short: "Create a snippet from a template",
	Long: `Create a snippet from a built-in template with interactive variable substitution.

Examples:
  snip insert go-func "My Go Function"
  snip insert py-class "User Class"
  snip insert --list  # Show available templates`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if insertList {
			templateList := templates.ListTemplates()
			fmt.Println("Available Templates:")
			fmt.Println("─────────────────────")
			for _, t := range templateList {
				fmt.Printf("• %s - %s (%s)\n", t.Name, t.Description, t.Language)
			}
			return nil
		}

		if len(args) < 2 {
			return fmt.Errorf("usage: snip insert [template-name] [title]")
		}

		templateName := args[0]
		title := args[1]

		// Get the template
		builtinTemplates := templates.GetBuiltinTemplates()
		template, exists := builtinTemplates[templateName]
		if !exists {
			fmt.Printf("Template '%s' not found. Available templates:\n", templateName)
			for name := range builtinTemplates {
				fmt.Printf("  - %s\n", name)
			}
			return nil
		}

		// Interactive variable collection
		vars := make(map[string]string)
		scanner := bufio.NewScanner(os.Stdin)

		// Extract variables from template content
		content := template.Content
		placeholders := extractPlaceholders(content)

		if len(placeholders) > 0 {
			fmt.Printf("Creating snippet from template '%s'\n", templateName)
			fmt.Println("Please provide values for the following variables:")
			fmt.Println()

			for _, placeholder := range placeholders {
				fmt.Printf("%s: ", placeholder)
				if scanner.Scan() {
					vars[placeholder] = scanner.Text()
				}
			}
		}

		// Create snippet from template
		snippet, err := templates.CreateSnippetFromTemplate(templateName, title, vars)
		if err != nil {
			return err
		}

		snippet.ID = uuid.NewString()

		// Save the snippet
		if err := snippets.SaveSnippet(*snippet); err != nil {
			return err
		}

		fmt.Printf("✅ Created snippet '%s' from template '%s'\n", title, templateName)
		fmt.Printf("   ID: %s\n", snippet.ID)
		return nil
	},
}

// extractPlaceholders extracts {{.Variable}} placeholders from template content
func extractPlaceholders(content string) []string {
	var placeholders []string
	seen := make(map[string]bool)

	// Simple regex-like extraction
	for {
		start := strings.Index(content, "{{.")
		if start == -1 {
			break
		}
		end := strings.Index(content[start:], "}}")
		if end == -1 {
			break
		}
		
		placeholder := content[start+3 : start+end]
		if !seen[placeholder] {
			placeholders = append(placeholders, placeholder)
			seen[placeholder] = true
		}
		
		content = content[start+end+2:]
	}

	return placeholders
}

func init() {
	rootCmd.AddCommand(insertCmd)
	insertCmd.Flags().BoolVarP(&insertList, "list", "l", false, "List available templates")
}
