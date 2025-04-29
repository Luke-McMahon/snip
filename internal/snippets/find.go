/*
Copyright © 2025 Luke McMahon <me@lmc.id.au>
*/
package snippets

import "os"

func FindSnippetByID(id string) (*Snippet, error) {
   all, err := LoadSnippets()
if err != nil { return nil, err }

  for _, s := range all {
	if s.ID == id {
		return &s, nil
	}
  }

  return nil, os.ErrNotExist
}
