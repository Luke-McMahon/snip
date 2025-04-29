/*
Copyright © 2025 Luke McMahon <me@lmc.id.au>
*/
package snippets

import "time"

type Snippet struct {
    ID        string    `json:"id"`
    Title     string    `json:"title"`
    Tags      []string  `json:"tags"`
    Content   string    `json:"content"`
    Language  string    `json:"language"`
    Starred   bool      `json:"starred"`
    Private   bool      `json:"private"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

