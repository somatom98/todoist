package todo

import "context"

type Repo interface {
	GetAll(ctx context.Context) ([]*Item, error)
	Get(ctx context.Context, collection string) ([]*Item, error)
	Collections(ctx context.Context) ([]Collection, error)
	Add(ctx context.Context, item *Item) error
}
