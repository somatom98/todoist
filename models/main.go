package models

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/somatom98/todoist/todo"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type mainModel struct {
	todoRepo    todo.Repo
	currentView int
	models      []tea.Model
}

func NewMain(todoRepo todo.Repo) *mainModel {
	return &mainModel{
		todoRepo:    todoRepo,
		currentView: int(collectionSelectorModel),
		models: []tea.Model{
			collectionSelectorModel: NewCollectionSelector(todoRepo),
			todoListModel:           NewTodoList(todoRepo),
		},
	}
}

func (m *mainModel) Init() tea.Cmd {
	return todo.UpdateCmd(todo.UpdateMsg{})
}

func (m *mainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	log.Printf("MAIN, msg: %T - %+v", msg, msg)
	cmds := []tea.Cmd{}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "tab":
			m.currentView = (m.currentView + 1) % len(m.models)
		default:
			dest, message := mapMessage(model(m.currentView), msg)

			var cmd tea.Cmd
			m.models[dest], cmd = m.models[dest].Update(message)
			cmds = append(cmds, cmd)
		}
	default:
		for i, mod := range m.models {
			var cmd tea.Cmd
			m.models[i], cmd = mod.Update(msg)
			cmds = append(cmds, cmd)
		}
	}

	return m, tea.Batch(cmds...)
}

func (m *mainModel) View() string {
	var s string
	renders := []string{}
	for _, mod := range m.models {
		renders = append(renders, docStyle.Render(mod.View()))
	}
	s += lipgloss.JoinHorizontal(lipgloss.Top, renders...)
	return s
}
