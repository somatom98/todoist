package todo

import (
	"context"
	"fmt"
)

type mockRepo struct {
	todos []*Todo
}

var _ TodoRepo = &mockRepo{}

func NewMockRepo() *mockRepo {
	return &mockRepo{
		todos: []*Todo{
			{
				title:       "First todo",
				description: "Description",
				done:        false,
				collection:  "main",
			},
			{
				title:       "Second todo",
				description: "Description",
				done:        true,
				collection:  "secondary",
			},
			{
				title:       "Second todo",
				description: "Description",
				done:        true,
				collection:  "main",
			},
		},
	}
}

func (r *mockRepo) GetAll(ctx context.Context) ([]*Todo, error) {
	return r.todos, nil
}

func (r *mockRepo) Get(ctx context.Context, collection string) ([]*Todo, error) {
	filtered := []*Todo{}

	for _, item := range r.todos {
		if item.collection == collection {
			filtered = append(filtered, item)
		}
	}

	return filtered, nil
}

func (r *mockRepo) Add(ctx context.Context, newItem *Todo) error {
	for _, item := range r.todos {
		if newItem.collection == item.collection && newItem.title == item.title {
			return fmt.Errorf("duplicate item")
		}
	}
	r.todos = append(r.todos, newItem)
	return nil
}
