# 🗃️ snipit 

A fast, local-first CLI snippet manager written in Go. Store, tag, and retrieve code snippets straight from your terminal — optionally with a TUI interface coming soon.

---

## 🚀 Features

- Add code snippets with title, tags, language, and metadata
- Open your `$EDITOR` to write snippets interactively
- Search or list all saved snippets
- Local JSON-based storage (no cloud or sync by default)
- Designed to be easily extended with TUI (via [Bubble Tea](https://github.com/charmbracelet/bubbletea))

---

## 📦 Installation

```sh
git clone https://github.com/luke-mcmahon/snipit.git
cd snipvault
go build -o snip
./snip
