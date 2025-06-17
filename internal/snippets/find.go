/*
Copyright © 2025 Luke McMahon <me@lmc.id.au>
*/
package snippets

import (
	"fmt"
	"os"
	"strings"
)

func FindSnippetByID(id string) (*Snippet, error) {
	all, err := LoadSnippets()
	if err != nil {
		return nil, err
	}

	// First try exact match
	for _, s := range all {
		if s.ID == id {
			return &s, nil
		}
	}

	// Then try partial match
	var matches []Snippet
	for _, s := range all {
		if strings.HasPrefix(s.ID, id) {
			matches = append(matches, s)
		}
	}

	if len(matches) == 1 {
		return &matches[0], nil
	} else if len(matches) > 1 {
		return nil, fmt.Errorf("ambiguous ID '%s' matches %d snippets", id, len(matches))
	}

	return nil, os.ErrNotExist
}
