package main

import (
	"context"
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/somatom98/todoist/todo"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type model struct {
	todoRepo todo.TodoRepo
	choices  []string
	list     list.Model
	cursor   int
	selected map[int]struct{}
}

func newModel() *model {
	return &model{
		todoRepo: todo.NewMockRepo(),
		choices:  []string{"New"},
		cursor:   0,
		selected: make(map[int]struct{}),
	}
}

func (m *model) Init() tea.Cmd {
	return m.getTodoCommand
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "a":
			err := m.todoRepo.Add(context.TODO(), todo.New("NEW", "TODO"))
			if err != nil {
				// TODO: popup
				break
			}
			m.list.InsertItem(0, todo.New("NEW", "TODO"))
		case " ", "enter":
			item := m.list.SelectedItem().(*todo.Todo)
			item.ChangeStatus()
			m.list.SetItem(m.list.Index(), item)
		}
	case getTodoCommandResponse:
		items := []list.Item{}
		for _, todo := range msg.todos {
			items = append(items, todo)
		}
		m.list = list.New(items, list.NewDefaultDelegate(), 0, 0)
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m *model) View() string {
	return docStyle.Render(m.list.View())
}

func main() {
	p := tea.NewProgram(newModel())

	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
