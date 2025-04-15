package snippets

import (
    "encoding/json"
    "os"
)

func LoadSnippets() ([]Snippet, error) {
    data, err := os.ReadFile(storagePath)
    if err != nil {
        return nil, err
    }

    var snippets []Snippet
    err = json.Unmarshal(data, &snippets)
    return snippets, err
}
