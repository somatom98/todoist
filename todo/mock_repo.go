package todo

import "context"

type mockRepo struct {
	todos []Todo
}

func NewMockRepo() *mockRepo {
	return &mockRepo{
		todos: []Todo{
			{
				title: "First todo",
			},
			{
				title: "Second todo",
			},
		},
	}
}

func (r *mockRepo) GetAll(ctx context.Context) ([]Todo, error) {
	return r.todos, nil
}

func (r *mockRepo) Add(ctx context.Context, item Todo) error {
	r.todos = append(r.todos, item)
	return nil
}
