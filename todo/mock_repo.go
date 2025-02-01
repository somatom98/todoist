package todo

import (
	"context"
	"fmt"
)

type mockRepo struct {
	todos []*Todo
}

func NewMockRepo() *mockRepo {
	return &mockRepo{
		todos: []*Todo{
			{
				title:       "First todo",
				description: "Description",
				done:        false,
			},
			{
				title:       "Second todo",
				description: "Description",
				done:        true,
			},
		},
	}
}

func (r *mockRepo) GetAll(ctx context.Context) ([]*Todo, error) {
	return r.todos, nil
}

func (r *mockRepo) Add(ctx context.Context, newItem *Todo) error {
	for _, item := range r.todos {
		if newItem.title == item.title {
			return fmt.Errorf("duplicate item")
		}
	}
	r.todos = append(r.todos, newItem)
	return nil
}
