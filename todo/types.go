package todo

import "context"

type Repo interface {
	Get(ctx context.Context, collection Collection, status Status) ([]Item, error)
	Collections(ctx context.Context) ([]Collection, error)
	Add(ctx context.Context, item Item) error
	Update(ctx context.Context, id int64, item Item) error
}
