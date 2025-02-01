package collectionselector

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
	ctx      context.Context
	todoRepo todo.TodoRepo
	list     list.Model
}

func New() *model {
	return &model{
		ctx:      context.Background(),
		todoRepo: todo.NewMockRepo(),
	}
}

func (m *model) Init() tea.Cmd {
	log.Printf("SELECTOR, init")
	return getCollectionsCmd
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	log.Printf("SELECTOR, msg: %T - %+v", msg, msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "a":
			item := todo.New("NEW", "TODO", "new")
			err := m.todoRepo.Add(context.TODO(), item)
			if err != nil {
				// TODO: popup
				break
			}
			// TODO: update items list (and maybe change focus on other view?)
		}
	case getCollectionsMsg:
		collections, err := m.todoRepo.Collections(m.ctx)
		if err != nil {
			// TODO: popup
			break
		}

		items := []list.Item{}
		for _, todo := range collections {
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
