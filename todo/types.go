package todo

import (
	"context"

	"github.com/charmbracelet/bubbles/list"
)

type Todo struct {
	title       string
	description string
}

var _ list.Item = Todo{}

func New(title, description string) Todo {
	return Todo{
		title:       title,
		description: description,
	}
}

func (t Todo) String() string {
	return t.title
}

func (t Todo) FilterValue() string {
	return t.title
}

func (t Todo) Title() string {
	return t.title
}

func (t Todo) Description() string {
	return t.description
}

type TodoRepo interface {
	GetAll(ctx context.Context) ([]Todo, error)
	Add(ctx context.Context, item Todo) error
}
