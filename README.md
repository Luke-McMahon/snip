# 🗃️ snip

A fast, local-first CLI snippet manager written in Go.  
Store, tag, and retrieve code snippets straight from your terminal — with an optional TUI interface coming soon.

---

## 🚀 Features

- Add code snippets with title, tags, language, and metadata
- Open your `$EDITOR` to write snippets interactively
- Syntax-highlight snippets automatically (using Chroma)
- List, search, and view saved snippets
- Local JSON-based storage (no cloud or sync by default)
- Designed to be easily extended with a TUI (via [Bubble Tea](https://github.com/charmbracelet/bubbletea))

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

## 🔭 Future Work

- `delete [id]` — remove a snippet
- `star` / `unstar` — mark favorite snippets
- `export` / `import` — backup and restore, possibly with JSON/YAML
- GitHub Gist sync or cloud export (optional, opt-in)
- `snip stats` — see top tags/languages used
- Templated snippet creation (`snip insert <template>`)
- Interactive TUI mode with:
  - Search + filter by tag/language
  - Preview + view + copy to clipboard
  - Batch operations (delete, export)
- Auto language detection (fallback when `--language` is missing)
- Plugin/hooks system for custom snippet actions

