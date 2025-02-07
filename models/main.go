package models

import (
	"context"
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/somatom98/todoist/todo"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type mainModel struct {
	todoRepo    todo.Repo
	focusedView int
	models      []tea.Model
	collection  todo.Collection
}

func NewMain(todoRepo todo.Repo) *mainModel {
	return &mainModel{
		todoRepo:    todoRepo,
		focusedView: int(viewCollectionSelector),
		models: []tea.Model{
			viewCollectionSelector: NewCollectionSelector(todoRepo),
			viewTodoList:           NewTodoList(todoRepo),
			viewItemForm:           NewItemFormModel(),
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
		if m.focusedView == int(viewItemForm) {
			var cmd tea.Cmd
			m.models[viewItemForm], cmd = m.models[viewItemForm].Update(msg)
			cmds = append(cmds, cmd)
			break
		}
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "tab":
			m.changeFocusedView()
		case "a":
			cmds = append(cmds, ViewCmd(ViewMsg{View: viewItemForm}))
		default:
			dest, message := mapMessage(view(m.focusedView), msg)

			var cmd tea.Cmd
			m.models[dest], cmd = m.models[dest].Update(message)
			cmds = append(cmds, cmd)
		}
	case ViewMsg:
		if msg.View == viewTodoList {
			var cmd tea.Cmd
			m.models[viewItemForm], cmd = m.models[viewItemForm].Update(msg)
			cmds = append(cmds, cmd)
		}
		m.focusedView = int(msg.View)
	case todo.AddMsg:
		err := m.todoRepo.Add(context.TODO(), todo.Item(msg))
		if err != nil {
			// TODO: popup
			log.Printf("err: %w", err)
			break
		}
		return m, todo.UpdateCmd(todo.UpdateMsg{})
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

	log.Printf("focused view: %v", m.focusedView)

	switch m.focusedView {
	case int(viewTodoList), int(viewCollectionSelector):
		renders = append(renders, docStyle.Render(m.models[viewCollectionSelector].View()))
		renders = append(renders, docStyle.Render(m.models[viewTodoList].View()))
	case int(viewItemForm):
		renders = append(renders, docStyle.Render(m.models[viewItemForm].View()))
	}

	s += lipgloss.JoinHorizontal(lipgloss.Top, renders...)
	return s
}

func (m *mainModel) changeFocusedView() {
	switch m.focusedView {
	case int(viewCollectionSelector):
		m.focusedView = int(viewTodoList)
	case int(viewTodoList):
		m.focusedView = int(viewCollectionSelector)
	}
}
