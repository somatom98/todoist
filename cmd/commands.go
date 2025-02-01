package main

import (
	"context"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/somatom98/todoist/todo"
)

type getTodoCommandResponse struct {
	todos []*todo.Todo
}

func (m *model) getTodoCommand() tea.Msg {
	ctx := context.Background()

	todos, err := m.todoRepo.Get(ctx, m.collection)
	if err != nil {
		panic("unable to get todo list")
	}

	return getTodoCommandResponse{todos}
}
