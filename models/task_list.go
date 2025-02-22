package models

import (
	"context"
	"log"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/somatom98/todoist/todo"
)

type taskList struct {
	status     todo.Status
	ctx        context.Context
	todoRepo   todo.Repo
	list       list.Model
	collection todo.Collection
	current    todo.Item
}

func NewTaskList(status todo.Status, todoRepo todo.Repo) *taskList {
	return &taskList{
		status:   status,
		ctx:      context.Background(),
		todoRepo: todoRepo,
		list:     list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0),
	}
}

func (m *taskList) Init() tea.Cmd {
	return todo.UpdateCmd(todo.UpdateMsg{})
}

func (m *taskList) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	log.Printf("TASKLIST, msg: %T - %+v", msg, msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case " ", "enter":
			item := m.list.SelectedItem().(todo.Item)
			item = item.UpdateStatus()
			m.todoRepo.Update(m.ctx, item.ID, item)
			return m, todo.UpdateCmd(todo.UpdateMsg{})
		case "a":
			return m, tea.Batch(
				todo.OperationCmd(todo.OperationAdd, m.current),
				ViewCmd(ViewMsg{View: viewItemForm}),
			)
		case "c":
			return m, tea.Batch(
				todo.OperationCmd(todo.OperationChange, m.current),
				ViewCmd(ViewMsg{View: viewItemForm}),
			)
		}
	case todo.UpdateMsg:
		if msg.Collection != nil {
			m.collection = *msg.Collection
		}

		todos, err := m.todoRepo.Get(m.ctx, m.collection, m.status)
		if err != nil {
			// TODO: popup
			log.Printf("err: %w", err)
			break
		}

		items := []list.Item{}
		for _, todo := range todos {
			items = append(items, todo)
		}
		m.list.SetItems(items)
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	item := m.list.SelectedItem()
	if item != nil && m.current != item.(todo.Item) {
		m.current = item.(todo.Item)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m *taskList) View() string {
	m.list.Title = string(m.status)
	return docStyle.Render(m.list.View())
}
