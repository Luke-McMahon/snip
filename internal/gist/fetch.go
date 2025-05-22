package gist

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/Luke-McMahon/snip/internal/snippets"
	"github.com/google/uuid"
)

// GistResponse represents the structure of a GitHub Gist API response
type GistResponse struct {
	ID          string                      `json:"id"`
	Description string                      `json:"description"`
	Public      bool                        `json:"public"`
	CreatedAt   string                      `json:"created_at"`
	UpdatedAt   string                      `json:"updated_at"`
	Files       map[string]GistFileResponse `json:"files"`
}

// GistFileResponse represents a file in a Gist response
type GistFileResponse struct {
	Filename string `json:"filename"`
	Type     string `json:"type"`
	Language string `json:"language"`
	Content  string `json:"content"`
}

// FetchUserGists fetches all gists for the authenticated user
// and returns them as a slice of Gists
func FetchUserGists() ([]GistResponse, error) {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		return nil, errors.New("GitHub token not found. Set the GITHUB_TOKEN environment variable")
	}

	// Create the HTTP request
	req, err := http.NewRequest("GET", githubAPIURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	// Set headers
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", "Bearer "+token)

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("GitHub API returned status code %d", resp.StatusCode)
	}

	// Read and parse the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}

	var gists []GistResponse
	if err := json.Unmarshal(body, &gists); err != nil {
		return nil, fmt.Errorf("error parsing JSON: %w", err)
	}

	return gists, nil
}

// ConvertGistToSnippet converts a GitHub Gist to a local Snippet
func ConvertGistToSnippet(gist GistResponse) (*snippets.Snippet, error) {
	if len(gist.Files) == 0 {
		return nil, errors.New("gist has no files")
	}

	// Just take the first file from the gist
	var file GistFileResponse
	for _, f := range gist.Files {
		file = f
		break
	}

	// Parse the created time
	createdAt, err := time.Parse(time.RFC3339, gist.CreatedAt)
	if err != nil {
		createdAt = time.Now()
	}

	// Parse the updated time
	updatedAt, err := time.Parse(time.RFC3339, gist.UpdatedAt)
	if err != nil {
		updatedAt = time.Now()
	}

	// Create a new snippet
	snippet := &snippets.Snippet{
		ID:        uuid.NewString(),
		Title:     gist.Description,
		Content:   file.Content,
		Language:  file.Language,
		Private:   !gist.Public,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Tags:      []string{"github", "gist", "imported"},
	}

	return snippet, nil
}

// ImportUserGists imports all user gists as local snippets
func ImportUserGists() error {
	// Fetch all gists
	gists, err := FetchUserGists()
	if err != nil {
		return err
	}

	// Load existing snippets
	allSnippets, err := snippets.LoadSnippets()
	if err != nil {
		return err
	}

	// Convert and add each gist
	imported := 0
	for _, gist := range gists {
		snippet, err := ConvertGistToSnippet(gist)
		if err != nil {
			// Skip this gist if there's an error
			continue
		}

		allSnippets = append(allSnippets, *snippet)
		imported++
	}

	// Save all snippets
	if err := snippets.SaveAllSnippets(allSnippets); err != nil {
		return err
	}

	return nil
}
