package todo

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

type Item struct {
	title       string
	description string
	done        bool
	collection  Collection
}

var _ list.Item = &Item{}

func New(title, description, collection string) *Item {
	return &Item{
		title:       title,
		description: description,
		done:        false,
	}
}

func (i *Item) ChangeStatus() {
	i.done = !i.done
}

func (i *Item) String() string {
	return i.title
}

// Implement list.Item interface

func (i *Item) FilterValue() string {
	return i.title
}

func (i *Item) Title() string {
	mark := "[ ]"
	if i.done {
		mark = lipgloss.NewStyle().
			Foreground(lipgloss.Color("2")).
			Render("[✔]")
	}
	return fmt.Sprintf("%s %s", mark, i.title)
}

func (i *Item) Description() string {
	return fmt.Sprintf("    %s", i.title)
}
