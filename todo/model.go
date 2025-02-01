package todo

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

type Todo struct {
	title       string
	description string
	done        bool
}

var _ list.Item = &Todo{}

func New(title, description string) *Todo {
	return &Todo{
		title:       title,
		description: description,
		done:        false,
	}
}

func (t *Todo) ChangeStatus() {
	t.done = !t.done
}

func (t *Todo) String() string {
	return t.title
}

// Implement list.Item interface

func (t *Todo) FilterValue() string {
	return t.title
}

func (t *Todo) Title() string {
	mark := "[ ]"
	if t.done {
		mark = lipgloss.NewStyle().
			Foreground(lipgloss.Color("2")).
			Render("[✔]")
	}
	return fmt.Sprintf("%s %s", mark, t.title)
}

func (t *Todo) Description() string {
	return fmt.Sprintf("    %s", t.title)
}
