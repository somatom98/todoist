package todo

import "context"

type Repo interface {
	Get(ctx context.Context, collection Collection) ([]*Item, error)
	Collections(ctx context.Context) ([]Collection, error)
	Add(ctx context.Context, item *Item) error
}
