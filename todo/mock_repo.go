package todo

import (
	"context"
	"fmt"
)

type mockRepo struct {
	items []Item
}

var _ Repo = &mockRepo{}

func NewMockRepo() *mockRepo {
	return &mockRepo{
		items: []Item{
			{
				Tit:        "First todo",
				Descr:      "Description",
				Status:     "todo",
				Collection: "main",
			},
			{
				Tit:        "Second todo",
				Descr:      "Description",
				Status:     "done",
				Collection: "secondary",
			},
			{
				Tit:        "Second todo",
				Descr:      "Description",
				Status:     "done",
				Collection: "main",
			},
		},
	}
}

func (r *mockRepo) Get(ctx context.Context, collection Collection, status Status) ([]Item, error) {
	filtered := []Item{}

	for _, item := range r.items {
		if item.Collection == collection && item.Status == status {
			filtered = append(filtered, item)
		}
	}

	return filtered, nil
}

func (r *mockRepo) Collections(ctx context.Context) ([]Collection, error) {
	checked := map[Collection]bool{}
	collections := []Collection{}

	for _, item := range r.items {
		if _, ok := checked[item.Collection]; !ok {
			checked[item.Collection] = true
			collections = append(collections, item.Collection)
		}
	}

	return collections, nil
}

func (r *mockRepo) Add(ctx context.Context, newItem Item) error {
	for _, item := range r.items {
		if newItem.Collection == item.Collection && newItem.Tit == item.Tit {
			return fmt.Errorf("duplicate item")
		}
	}
	r.items = append(r.items, newItem)
	return nil
}

func (r *mockRepo) Update(ctx context.Context, id int64, item Item) error {
	for i, it := range r.items {
		if it.Tit == item.Tit {
			r.items[i] = item
		}
	}
	return nil
}
