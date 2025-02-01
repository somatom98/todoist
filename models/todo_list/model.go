package todolist

import (
	"context"
	"log"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/somatom98/todoist/todo"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type model struct {
	todoRepo   todo.TodoRepo
	list       list.Model
	collection string
}

func New() *model {
	return &model{
		todoRepo:   todo.NewMockRepo(),
		collection: "main",
	}
}

func (m *model) Init() tea.Cmd {
	log.Printf("TODOLIST, init")
	return m.getItemCmd
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	log.Printf("TODOLIST, msg: %T - %+v", msg, msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "a":
			item := todo.New("NEW", "TODO", "main")
			err := m.todoRepo.Add(context.TODO(), item)
			if err != nil {
				// TODO: popup
				break
			}
			m.list.InsertItem(0, item)
		case " ", "enter":
			item := m.list.SelectedItem().(*todo.Item)
			item.ChangeStatus()
			m.list.SetItem(m.list.Index(), item)
		}
	case getItemsMsg:
		items := []list.Item{}
		for _, todo := range msg.items {
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
