package todo

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
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
	i.Status = i.Status.Next()
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
	return fmt.Sprintf("    %s", i.Descr)
}
