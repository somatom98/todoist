package todo

import "context"

type TodoRepo interface {
	GetAll(ctx context.Context) ([]*Todo, error)
	Get(ctx context.Context, collection string) ([]*Todo, error)
	Add(ctx context.Context, item *Todo) error
}
