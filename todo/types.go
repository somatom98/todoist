package todo

import (
	"context"

	"github.com/charmbracelet/bubbles/list"
)

type Todo struct {
	title string
}

var _ list.Item = Todo{}

func (t *Todo) String() string {
	return t.title
}

func (t Todo) FilterValue() string {
	return t.title
}

func (t Todo) Title() string {
	return t.title
}

func (t Todo) Description() string {
	return ""
}

type TodoRepo interface {
	GetAll(ctx context.Context) ([]Todo, error)
}
