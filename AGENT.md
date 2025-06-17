# AGENT.md - Snip CLI Development Guide

## Build/Lint/Test Commands
- **Build**: `go build -o snip`
- **Test**: `go test ./...` (no tests exist yet)
- **Test single package**: `go test ./internal/snippets`
- **Run locally**: `go run main.go` or `./snip`
- **Install**: `go install`

## Architecture
- **Entry point**: `main.go` → `cmd.Execute()`
- **Commands**: `cmd/` - Cobra CLI commands (add, list, search, view, edit, gist, import)
- **Core logic**: `internal/snippets/` - snippet CRUD operations, JSON storage
- **GitHub integration**: `internal/gist/` - GitHub Gist import/export
- **Storage**: Local JSON file at `$HOME/.snippets/snippets.json`

## Code Style
- **Naming**: CamelCase exports, camelCase internal, descriptive names
- **Errors**: Use `fmt.Errorf` with `%w` wrapping, early returns
- **JSON**: snake_case tags (`created_at`, `updated_at`)
- **Headers**: Copyright notice required on new files
- **Packages**: Single lowercase words, clear separation of concerns
- **Imports**: Standard → third-party → internal, grouped with blank lines
