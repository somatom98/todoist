package controllers

import (
	"context"
	"database/sql"

	"github.com/somatom98/todoist/db"
	"github.com/somatom98/todoist/domain"
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

func (r *repo) Add(ctx context.Context, item domain.Item) error {
	params := db.AddItemParams{
		Title:       item.Tit,
		Description: item.Descr,
		Collection:  string(item.Collection),
	}
	return r.db.AddItem(ctx, params)
}

func (r *repo) Collections(ctx context.Context) ([]domain.Collection, error) {
	coll, err := r.db.GetCollections(ctx)
	if err != nil {
		return nil, err
	}

	collections := []domain.Collection{}
	for _, collection := range coll {
		collections = append(collections, domain.Collection(collection))
	}

	return collections, nil
}

func (r *repo) Get(ctx context.Context, collection domain.Collection, status domain.Status) ([]domain.Item, error) {
	items, err := r.db.GetItems(ctx, db.GetItemsParams{
		Collection: string(collection),
		Status:     string(status),
	})
	if err != nil {
		return nil, err
	}

	mappedItems := []domain.Item{}
	for _, item := range items {
		mappedItems = append(mappedItems, domain.Item{
			ID:         item.ID,
			Tit:        item.Title,
			Descr:      item.Description,
			Status:     domain.Status(item.Status),
			Collection: domain.Collection(item.Collection),
		})
	}
	return mappedItems, nil
}

func (r *repo) Update(ctx context.Context, id int64, item domain.Item) error {
	return r.db.UpdateItem(ctx, db.UpdateItemParams{
		Title:       item.Tit,
		Description: item.Descr,
		Status:      string(item.Status),
		Collection:  string(item.Collection),
		ID:          id,
	})
}
