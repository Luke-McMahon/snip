# 🗃️ snip

A fast, local-first CLI snippet manager written in Go.  
Store, tag, and retrieve code snippets straight from your terminal — with an optional TUI interface coming soon.

---

## 🚀 Features

- **Add & Edit** - Create snippets with title, tags, language, and metadata
- **Smart Language Detection** - Auto-detect programming language from content
- **Interactive Editing** - Open your `$EDITOR` to write snippets interactively
- **Syntax Highlighting** - View snippets with automatic syntax highlighting (using Chroma)
- **Search & Filter** - List, search, and view saved snippets with filtering options
- **Star & Organize** - Mark favorite snippets and organize with tags
- **Statistics** - View usage stats including top tags and languages
- **Templates** - Create snippets from built-in templates (Go, Python, JS, etc.)
- **Copy to Clipboard** - Quickly copy snippet content to system clipboard
- **Backup & Restore** - Export/import snippets in JSON or YAML format
- **GitHub Integration** - Import from and export to GitHub Gists
- **Local Storage** - JSON-based storage with no cloud dependencies
- **Cross-Platform** - Works on macOS, Linux, and Windows

---

## 📦 Installation

### Using Go

```sh
go install github.com/luke-mcmahon/snip@latest
```

### From Source

```sh
git clone https://github.com/luke-mcmahon/snip.git
cd snip
go build -o snip
./snip
```

---

## 🛠️ Usage

### End to end

```sh
# Add a new snippet (title is positional)
snip add "Reverse a string in Go" \
  --tags go,string,utility \
  --language go

# Language detection is automatic if not specified!
snip add "Quick Python Script" --tags python,utility

# You'll be prompted to enter the snippet content in your default editor (e.g. nano, vim)
# Example content:
# func Reverse(s string) string {
#     runes := []rune(s)
#     for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
#         runes[i], runes[j] = runes[j], runes[i]
#     }
#     return string(runes)
# }

# List all snippets
snip list

# Output:
# - Reverse a string in Go [go, string, utility]
#         ID: b5c0478e-379c-43ab-9f06-a01711fe54b0

# Search snippets by tag
snip search --tag go

# Search snippets by keyword
snip search reverse

# View the full snippet
snip view b5c0478e-379c-43ab-9f06-a01711fe54b0

# Output:
# ────────────────────────────
# 📌 Reverse a string in Go [go, string, utility]
# 🆔 b5c0478e-379c-43ab-9f06-a01711fe54b0
# 🗣️  Language: go
# 🕒 Created: 2025-04-29 21:09
# 🕓 Updated: 2025-04-29 21:09
# ────────────────────────────
# func Reverse(s string) string {
#     runes := []rune(s)
#     for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
#         runes[i], runes[j] = runes[j], runes[i]
#     }
#     return string(runes)
# }
# ────────────────────────────
```

### Add a snippet:

```sh
snip add "Curl JSON POST" --tags http,curl --language bash
```

If `--content` is not provided, your `$EDITOR` will open for you to complete.

#### Inline addition

```sh
snip add "Hello World in Python" \
  --language python \
  --tags hello,python \
  --content 'print("Hello, world!")'
```

---

### List all Snips

```sh
snip list
```

---

### Search snips

```sh
snip search curl
snip search json --tag go
snip search handler --lang go
```

---

### View a snip

```sh
snip view <id>
```

- Auto syntax-highlighting based on `language` applied to the snip
- Add `--no-highlight` if you want plain text output

```sh
snip view <id> --no-highlight
```

---

### Edit a snip

```sh
snip edit <id>
```

Opens the snippet content in your `$EDITOR` for quick updates.

---

### Star/Unstar snippets

```sh
snip star <id>      # Mark as favorite
snip unstar <id>    # Remove from favorites
```

### Filter lists

```sh
snip list --starred  # Show only starred snippets
snip list --private  # Show only private snippets
```

### Copy to clipboard

```sh
snip copy <id>  # Copy snippet content to clipboard
```

### Delete snippets

```sh
snip delete <id>  # Remove a snippet
```

### View statistics

```sh
snip stats  # Show snippet statistics, top tags, languages
```

### Templates

```sh
snip insert --list                    # List available templates
snip insert go-func "My Function"     # Create from Go function template
snip insert py-class "User Class"     # Create from Python class template
```

Available templates: `go-func`, `go-struct`, `py-func`, `py-class`, `js-func`, `js-arrow`, `bash-script`, `dockerfile`

### Backup & Restore

```sh
# Export all snippets
snip export                           # Export to JSON (default)
snip export --format yaml            # Export to YAML
snip export --output my-backup.json  # Custom filename

# Import snippets
snip import backup.json               # Import from JSON
snip import backup.yaml               # Import from YAML
snip import --merge=false backup.json  # Replace instead of merge
```

---

Each snippet includes:

- ID – unique identifier
- Title – short descriptive title
- Content – the actual code/text
- Tags – comma-separated categories
- Language – used for syntax highlighting
- Starred / Private – optional flags
- CreatedAt / UpdatedAt – timestamps

Stored locally in `$HOME/.snippets/snippets.json`

---

### Export to GitHub Gist

Export any snippet to a GitHub Gist:

```sh
# First, set your GitHub token in your environment
export GITHUB_TOKEN=your_github_personal_access_token

# Export a snippet to GitHub Gist
snip gist <snippet-id>
```

Make sure your GitHub token has the `gist` scope permission to create gists.

### Import from GitHub Gists

Import all your GitHub Gists as local snippets:

```sh
# First, set your GitHub token in your environment
export GITHUB_TOKEN=your_github_personal_access_token

# Import all your gists as snippets
snip import-gists
```

Imported snippets will automatically be tagged with `github`, `gist`, and `imported`.

---

## 🔭 Future Work

- ✅ `delete [id]` — remove a snippet *(implemented)*
- ✅ `star` / `unstar` — mark favorite snippets *(implemented)*
- ✅ `export` / `import` — backup and restore with JSON/YAML *(implemented)*
- ✅ `snip stats` — see top tags/languages used *(implemented)*
- ✅ Templated snippet creation (`snip insert <template>`) *(implemented)*
- ✅ Auto language detection (fallback when `--language` is missing) *(implemented)*
- ✅ Copy to clipboard functionality *(implemented)*
- ✅ GitHub Gist sync or cloud export *(implemented)*

**Remaining Future Work:**
- Interactive TUI mode with:
  - Search + filter by tag/language
  - Preview + view + copy to clipboard
  - Batch operations (delete, export)
- Plugin/hooks system for custom snippet actions
- Fuzzy search improvements
- Snippet sharing via URLs
- Version control integration
