package models

import (
	"context"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/somatom98/todoist/controllers"
	"github.com/somatom98/todoist/domain"
)

type collectionSelector struct {
	ctx      context.Context
	todoRepo controllers.ItemsRepo
	list     list.Model
	current  domain.Collection
	focused  bool
	width    int
	height   int
}

func NewCollectionSelector(todoRepo controllers.ItemsRepo) *collectionSelector {
	return &collectionSelector{
		ctx:      context.Background(),
		todoRepo: todoRepo,
		list:     list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0),
	}
}

func (m *collectionSelector) Init() tea.Cmd {
	return domain.UpdateCmd(domain.UpdateMsg{})
}

func (m *collectionSelector) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case domain.UpdateMsg:
		collections, err := m.todoRepo.Collections(m.ctx)
		if err != nil {
			// TODO: popup
			break
		}

		items := []list.Item{}
		for _, collection := range collections {
			items = append(items, collection)
		}
		m.list.SetItems(items)
	case tea.WindowSizeMsg:
		h, v := paneStyle.GetFrameSize()
		m.width = (msg.Width - h) / 4
		m.height = (msg.Height - v)
		m.list.SetSize(m.width, m.height)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)

	item := m.list.SelectedItem()
	if item != nil && m.current != item.(domain.Collection) {
		m.current = item.(domain.Collection)
		return m, tea.Batch(cmd, domain.UpdateCmd(domain.UpdateMsg{Collection: &m.current}))
	}

	return m, cmd
}

func (m *collectionSelector) View() string {
	style := paneStyle
	if m.focused {
		style = focusedPaneStyle
	}
	return style.Render(m.list.View())
}

func (m *collectionSelector) ChangeFocus() {
	m.focused = !m.focused
}
