package todo

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

type Status string

const (
	StatusTodo       = "todo"
	StatusInProgress = "in_progress"
	StatusDone       = "done"
)

type Item struct {
	ID         int64
	Tit        string
	Descr      string
	Status     Status
	Collection Collection
}

var _ list.Item = Item{}

func New(title, description, collection string) Item {
	return Item{
		Tit:        title,
		Descr:      description,
		Status:     "todo",
		Collection: Collection(collection),
	}
}

func (i Item) UpdateStatus() Item {
	switch i.Status {
	case StatusTodo:
		i.Status = StatusInProgress
	case StatusInProgress:
		i.Status = StatusDone
	}
	return i
}

// Implement list.Item interface

func (i Item) FilterValue() string {
	return i.Tit
}

func (i Item) Title() string {
	mark := "[ ]"
	if i.Status == StatusDone {
		mark = lipgloss.NewStyle().
			Foreground(lipgloss.Color("2")).
			Render("[✔]")
	}
	return fmt.Sprintf("%s %s", mark, i.Tit)
}

func (i Item) Description() string {
	return fmt.Sprintf("    %s", i.Tit)
}
