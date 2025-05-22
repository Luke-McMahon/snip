package gist

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/Luke-McMahon/snip/internal/snippets"
)

const (
	githubAPIURL = "https://api.github.com/gists"
)

// Save saves a snippet as a GitHub Gist and returns any error
func Save(snippet *snippets.Snippet) error {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		return errors.New("GitHub token not found. Set the GITHUB_TOKEN environment variable")
	}

	// Prepare the file content for the gist
	files := map[string]map[string]string{
		getFilename(snippet): {
			"content": snippet.Content,
		},
	}

	// Create the request body
	reqBody := map[string]interface{}{
		"description": snippet.Title,
		"public":      !snippet.Private,
		"files":       files,
	}

	// Convert to JSON
	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return fmt.Errorf("error marshaling JSON: %w", err)
	}

	// Create the HTTP request
	req, err := http.NewRequest("POST", githubAPIURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	// Set headers
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("GitHub API returned status code %d", resp.StatusCode)
	}

	fmt.Println("Snippet saved to GitHub Gist")
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %w", err)
	}

	var gist map[string]interface{}
	err = json.Unmarshal(body, &gist)
	if err != nil {
		return fmt.Errorf("error unmarshalling JSON: %w", err)
	}
	fmt.Println(gist["html_url"])

	return nil
}

// getFilename generates a filename for the gist based on snippet properties
func getFilename(snippet *snippets.Snippet) string {
	filename := snippet.Title

	// Add extension based on language if available
	if snippet.Language != "" {
		// Clean up language name and map to file extension
		lang := strings.ToLower(snippet.Language)
		switch lang {
		case "javascript", "js":
			filename += ".js"
		case "typescript", "ts":
			filename += ".ts"
		case "python", "py":
			filename += ".py"
		case "go", "golang":
			filename += ".go"
		case "ruby", "rb":
			filename += ".rb"
		case "java":
			filename += ".java"
		case "php":
			filename += ".php"
		case "css":
			filename += ".css"
		case "html":
			filename += ".html"
		case "markdown", "md":
			filename += ".md"
		case "shell", "bash", "sh":
			filename += ".sh"
		default:
			// Use language as extension if none of the above
			filename += "." + lang
		}
	} else {
		// Default to .txt if no language specified
		filename += ".txt"
	}

	return filename
}
