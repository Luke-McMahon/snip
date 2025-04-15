package snippets

import (
    "encoding/json"
	"fmt"
    "os"
    "path/filepath"
)

const storagePath = "snippets/snippets.json"

func SaveSnippet(s Snippet) error {
    var existing []Snippet

    // Ensure storage folder exists
    _ = os.MkdirAll(filepath.Dir(storagePath), 0755)

    // Read existing snippets if the file exists
    if data, err := os.ReadFile(storagePath); err == nil {
        _ = json.Unmarshal(data, &existing)
    }

    // Append new snippet
    existing = append(existing, s)

    // Marshal back to JSON
    updated, err := json.MarshalIndent(existing, "", "  ")
    if err != nil {
        return err
    }

	fmt.Printf("snipped it: %s\n", s.Title)

    // Save to disk
    return os.WriteFile(storagePath, updated, 0644)
}
