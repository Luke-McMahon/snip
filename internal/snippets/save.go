/*
Copyright © 2025 Luke McMahon <me@lmc.id.au>
*/
package snippets

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func getStoragePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(home, ".snippets", "snippets.json"), nil
}

func SaveSnippet(s Snippet) error {
	path, err := getStoragePath()
	if err != nil {
		return err
	}

	var existing []Snippet

	_ = os.MkdirAll(filepath.Dir(path), 0755)

	if data, err := os.ReadFile(path); err == nil {
		_ = json.Unmarshal(data, &existing)
	}

	existing = append(existing, s)

	updated, err := json.MarshalIndent(existing, "", "  ")
	if err != nil {
		return err
	}

	fmt.Printf("snipped it: %s\n", s.Title)

	return os.WriteFile(path, updated, 0644)
}

func SaveAllSnippets(snips []Snippet) error {
	updated, err := json.MarshalIndent(snips, "", "  ")
	if err != nil {
		return err
	}

	path, err := getStoragePath()
	if err != nil {
		return err
	}

	return os.WriteFile(path, updated, 0644)
}
