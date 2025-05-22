/*
Copyright © 2025 Luke McMahon <me@lmc.id.au>
*/
package snippets

import (
	"encoding/json"
	"os"
)

func LoadSnippets() ([]Snippet, error) {
	path, err := getStoragePath()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var snippets []Snippet
	err = json.Unmarshal(data, &snippets)
	return snippets, err
}
