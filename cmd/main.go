package main

import (
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/somatom98/todoist/todo"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type model struct {
	todoRepo    todo.TodoRepo
	currentView int
	models      []tea.Model
	collection  string
}

func newModel() *model {
	return &model{
		todoRepo:    todo.NewMockRepo(),
		currentView: 0,
		collection:  "main",
	}
}

func (m *model) Init() tea.Cmd {
	return initCommand
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			var cmd tea.Cmd
			m.models[m.currentView], cmd = m.models[m.currentView].Update(msg)
			cmds = append(cmds, cmd)
		}
	case initMsg:
		m.models = msg.Models
		for i, mod := range msg.Models {
			var cmd tea.Cmd
			initCmd := mod.Init()
			m.models[i], cmd = mod.Update(initCmd())
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

func (m *model) View() string {
	var s string
	renders := []string{}
	for _, mod := range m.models {
		renders = append(renders, docStyle.Render(mod.View()))
	}
	s += lipgloss.JoinHorizontal(lipgloss.Top, renders...)
	return s
}

func main() {
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		log.Fatalf("dead %w", err)
	}
	defer f.Close()

	p := tea.NewProgram(newModel())

	if _, err := p.Run(); err != nil {
		log.Fatalf("dead %w", err)
		os.Exit(1)
	}
}
