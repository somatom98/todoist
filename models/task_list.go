package models

import (
	"context"
	"log"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/somatom98/todoist/controllers"
	"github.com/somatom98/todoist/domain"
)

type taskList struct {
	status     domain.Status
	ctx        context.Context
	todoRepo   controllers.ItemsRepo
	list       list.Model
	collection domain.Collection
	current    domain.Item
	focused    bool
	width      int
	height     int
}

func NewTaskList(status domain.Status, todoRepo controllers.ItemsRepo) *taskList {
	return &taskList{
		status:   status,
		ctx:      context.Background(),
		todoRepo: todoRepo,
		list:     list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0),
	}
}

func (m *taskList) Init() tea.Cmd {
	return domain.UpdateCmd(domain.UpdateMsg{})
}

func (m *taskList) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case " ", "enter":
			item := m.list.SelectedItem().(domain.Item)
			item = item.UpdateStatus()
			m.todoRepo.Update(m.ctx, item.ID, item)
			return m, domain.UpdateCmd(domain.UpdateMsg{})
		case "a":
			return m, domain.OperationCmd(domain.OperationAdd, m.current)
		case "c":
			return m, domain.OperationCmd(domain.OperationChange, m.current)
		}
	case domain.UpdateMsg:
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
		h, v := paneStyle.GetFrameSize()
		m.width = (msg.Width - h) / 4
		m.height = msg.Height - v
		m.list.SetSize(m.width, m.height)
	}

	item := m.list.SelectedItem()
	if item != nil && m.current != item.(domain.Item) {
		m.current = item.(domain.Item)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m *taskList) View() string {
	m.list.Title = string(m.status)
	style := paneStyle
	if m.focused {
		style = focusedPaneStyle
	}
	return style.Width(m.width).Height(m.height).Render(m.list.View())
}

func (m *taskList) ChangeFocus() {
	m.focused = !m.focused
}
