package todo

import (
	"context"
	"database/sql"

	"github.com/somatom98/todoist/db"
)

type repo struct {
	db *db.Queries
}

var _ Repo = &repo{}

func NewRepo(database *sql.DB) *repo {
	return &repo{
		db: db.New(database),
	}
}

func (r *repo) Add(ctx context.Context, item Item) error {
	params := db.AddItemParams{
		Title:       item.Tit,
		Description: item.Descr,
		Collection:  string(item.Collection),
	}
	return r.db.AddItem(ctx, params)
}

func (r *repo) Collections(ctx context.Context) ([]Collection, error) {
	coll, err := r.db.GetCollections(ctx)
	if err != nil {
		return nil, err
	}

	collections := []Collection{}
	for _, collection := range coll {
		collections = append(collections, Collection(collection))
	}

	return collections, nil
}

func (r *repo) Get(ctx context.Context, collection Collection) ([]Item, error) {
	items, err := r.db.GetItems(ctx, string(collection))
	if err != nil {
		return nil, err
	}

	mappedItems := []Item{}
	for _, item := range items {
		mappedItems = append(mappedItems, Item{
			ID:         item.ID,
			Tit:        item.Title,
			Descr:      item.Description,
			Completed:  item.Completed,
			Collection: Collection(item.Collection),
		})
	}
	return mappedItems, nil
}

func (r *repo) Update(ctx context.Context, id int64, item Item) error {
	return r.db.UpdateItem(ctx, db.UpdateItemParams{
		Title:       item.Tit,
		Description: item.Descr,
		Completed:   item.Completed,
		Collection:  string(item.Collection),
		ID:          id,
	})
}
