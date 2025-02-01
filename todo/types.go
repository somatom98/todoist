package todo

import "context"

type TodoRepo interface {
	GetAll(ctx context.Context) ([]*Todo, error)
	Add(ctx context.Context, item *Todo) error
}
