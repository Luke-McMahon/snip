
// internal/tui/tui.go
package tui

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"

	"github.com/luke-mcmahon/snip/internal/snippets"
)

const (
	listWidth  = 50
	listHeight = 25
)

// snippetItem wraps a Snippet to satisfy list.Item
type snippetItem struct {
	snippets.Snippet
}

func (i snippetItem) Title() string {
	return i.Snippet.Title
}

func (i snippetItem) Description() string {
	return i.Snippet.Language
}

func (i snippetItem) FilterValue() string {
	return i.Snippet.Title
}

// model holds the Bubble Tea list
type model struct {
	list list.Model
}

// initialModel loads snippets and sets up the list
func initialModel() (model, error) {
	all, err := snippets.LoadSnippets()
	if err != nil {
		return model{}, fmt.Errorf("error loading snippets: %w", err)
	}
	items := make([]list.Item, len(all))
	for idx, sn := range all {
		items[idx] = snippetItem{sn}
	}
	l := list.New(items, list.NewDefaultDelegate(), listWidth, listHeight)
	// style the title
	titleStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("205"))
	l.Title = titleStyle.Render("Snipit Snippets")
	return model{list: l}, nil
}

// Init is called when the program starts
func (m model) Init() tea.Cmd {
	return tea.ClearScreen
}

// Update handles key events and delegates to the list
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

// View renders the list
func (m model) View() string {
	return m.list.View()
}

// Run initializes the model and starts the Bubble Tea program
func Run() error {
	m, err := initialModel()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	p := tea.NewProgram(m, tea.WithAltScreen())
	_, err = p.Run()
	if err != nil {
		return fmt.Errorf("error running TUI: %w", err)
	}
	return nil
}
