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
func FetchUserGists() ([]GistResponse, error) {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		return nil, errors.New("GitHub token not found. Set the GITHUB_TOKEN environment variable")
	}

	req, err := http.NewRequest("GET", githubAPIURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("GitHub API returned status code %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}

	var gistSummaries []GistResponse
	if err := json.Unmarshal(body, &gistSummaries); err != nil {
		return nil, fmt.Errorf("error parsing JSON: %w", err)
	}

	var fullGists []GistResponse
	for _, gistSummary := range gistSummaries {
		fullGist, err := FetchSingleGist(gistSummary.ID)
		if err != nil {
			fmt.Printf("Warning: Could not fetch gist %s: %v\n", gistSummary.ID, err)
			continue
		}
		fullGists = append(fullGists, fullGist)
	}

	return fullGists, nil
}

// FetchSingleGist fetches a single gist by its ID to get the complete content
func FetchSingleGist(gistID string) (GistResponse, error) {
	token := os.Getenv("GITHUB_TOKEN")
	singleGistURL := fmt.Sprintf("%s/%s", githubAPIURL, gistID)

	req, err := http.NewRequest("GET", singleGistURL, nil)
	if err != nil {
		return GistResponse{}, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return GistResponse{}, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return GistResponse{}, fmt.Errorf("GitHub API returned status code %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return GistResponse{}, fmt.Errorf("error reading response: %w", err)
	}

	var gist GistResponse
	if err := json.Unmarshal(body, &gist); err != nil {
		return GistResponse{}, fmt.Errorf("error parsing JSON: %w", err)
	}

	return gist, nil
}

// ConvertGistToSnippet converts a GitHub Gist to a local Snippet
func ConvertGistToSnippet(gist GistResponse) (*snippets.Snippet, error) {
	if len(gist.Files) == 0 {
		return nil, errors.New("gist has no files")
	}

	var file GistFileResponse
	for _, f := range gist.Files {
		file = f
		break
	}

	createdAt, err := time.Parse(time.RFC3339, gist.CreatedAt)
	if err != nil {
		createdAt = time.Now()
	}

	updatedAt, err := time.Parse(time.RFC3339, gist.UpdatedAt)
	if err != nil {
		updatedAt = time.Now()
	}

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
	gists, err := FetchUserGists()
	if err != nil {
		return err
	}

	allSnippets, err := snippets.LoadSnippets()
	if err != nil {
		return err
	}

	imported := 0
	for _, gist := range gists {
		snippet, err := ConvertGistToSnippet(gist)
		if err != nil {
			continue
		}

		allSnippets = append(allSnippets, *snippet)
		imported++
	}

	if err := snippets.SaveAllSnippets(allSnippets); err != nil {
		return err
	}

	return nil
}
