package todolist

import (
	"context"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/somatom98/todoist/todo"
)

type initMsg struct{}

func InitCmd() tea.Msg {
	return getItemsMsg{}
}

type getItemsMsg struct {
	items []*todo.Item
}

func (m *model) getItemCmd() tea.Msg {
	ctx := context.Background()

	todos, err := m.todoRepo.Get(ctx, m.collection)
	if err != nil {
		panic("unable to get todo list")
	}

	return getItemsMsg{todos}
}
