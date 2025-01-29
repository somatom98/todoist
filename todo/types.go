package todo

import "context"

type Todo struct {
	title string
}

func (t *Todo) String() string {
	return t.title
}

type TodoRepo interface {
	GetAll(ctx context.Context) ([]Todo, error)
}
