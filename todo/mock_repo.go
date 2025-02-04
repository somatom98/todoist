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
				title:       "First todo",
				description: "Description",
				completed:   false,
				collection:  "main",
			},
			{
				title:       "Second todo",
				description: "Description",
				completed:   true,
				collection:  "secondary",
			},
			{
				title:       "Second todo",
				description: "Description",
				completed:   true,
				collection:  "main",
			},
		},
	}
}

func (r *mockRepo) Get(ctx context.Context, collection Collection) ([]Item, error) {
	filtered := []Item{}

	for _, item := range r.items {
		if item.collection == collection {
			filtered = append(filtered, item)
		}
	}

	return filtered, nil
}

func (r *mockRepo) Collections(ctx context.Context) ([]Collection, error) {
	checked := map[Collection]bool{}
	collections := []Collection{}

	for _, item := range r.items {
		if _, ok := checked[item.collection]; !ok {
			checked[item.collection] = true
			collections = append(collections, item.collection)
		}
	}

	return collections, nil
}

func (r *mockRepo) Add(ctx context.Context, newItem Item) error {
	for _, item := range r.items {
		if newItem.collection == item.collection && newItem.title == item.title {
			return fmt.Errorf("duplicate item")
		}
	}
	r.items = append(r.items, newItem)
	return nil
}

func (r *mockRepo) Update(ctx context.Context, id int64, item Item) error {
	for i, it := range r.items {
		if it.title == item.title {
			r.items[i] = item
		}
	}
	return nil
}
