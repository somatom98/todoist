package models

import (
	"context"
	"log"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/somatom98/todoist/todo"
)

type collectionSelector struct {
	ctx      context.Context
	todoRepo todo.Repo
	list     list.Model
	current  todo.Collection
}

func NewCollectionSelector(todoRepo todo.Repo) *collectionSelector {
	return &collectionSelector{
		ctx:      context.Background(),
		todoRepo: todoRepo,
		list:     list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0),
	}
}

func (m *collectionSelector) Init() tea.Cmd {
	return todo.UpdateCmd(todo.UpdateMsg{})
}

func (m *collectionSelector) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	log.Printf("SELECTOR, msg: %T - %+v", msg, msg)

	switch msg := msg.(type) {
	case todo.UpdateMsg:
		collections, err := m.todoRepo.Collections(m.ctx)
		if err != nil {
			// TODO: popup
			log.Printf("err: %w", err)
			break
		}

		items := []list.Item{}
		for _, collection := range collections {
			items = append(items, collection)
		}
		m.list.SetItems(items)
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)

	item := m.list.SelectedItem()
	if item != nil && m.current != item.(todo.Collection) {
		m.current = item.(todo.Collection)
		return m, tea.Batch(cmd, todo.UpdateCmd(todo.UpdateMsg{Collection: &m.current}))
	}

	return m, cmd
}

func (m *collectionSelector) View() string {
	return docStyle.Render(m.list.View())
}
