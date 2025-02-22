package models

import (
	"context"
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/somatom98/todoist/controllers"
	"github.com/somatom98/todoist/domain"
)

var (
	docStyle         = lipgloss.NewStyle().Margin(1, 2)
	focusedViewStyle = lipgloss.NewStyle().Border(lipgloss.ThickBorder())
)

type mainModel struct {
	todoRepo    controllers.ItemsRepo
	focusedView int
	models      []tea.Model
	collection  domain.Collection
}

func NewMain(todoRepo controllers.ItemsRepo) *mainModel {
	return &mainModel{
		todoRepo:    todoRepo,
		focusedView: int(domain.ViewCollectionSelector),
		models: []tea.Model{
			domain.ViewCollectionSelector: NewCollectionSelector(todoRepo),
			domain.ViewTodoList:           NewTaskList(domain.StatusTodo, todoRepo),
			domain.ViewInProgressList:     NewTaskList(domain.StatusInProgress, todoRepo),
			domain.ViewDoneList:           NewTaskList(domain.StatusDone, todoRepo),
			domain.ViewItemForm:           NewItemFormModel(),
		},
	}
}

func (m *mainModel) Init() tea.Cmd {
	return domain.UpdateCmd(domain.UpdateMsg{})
}

func (m *mainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	log.Printf("MAIN, msg: %T - %+v", msg, msg)
	cmds := []tea.Cmd{}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.focusedView == int(domain.ViewItemForm) {
			cmds = append(cmds, m.update(domain.ViewItemForm, msg))
			break
		}
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "tab":
			m.changeFocusedView()
		default:
			dest, message := mapMessage(domain.View(m.focusedView), msg)
			cmds = append(cmds, m.update(dest, message))
		}
	case ViewMsg:
		m.focusedView = int(msg.View)
	case domain.AddMsg:
		err := m.todoRepo.Add(context.TODO(), domain.Item(msg))
		if err != nil {
			// TODO: popup
			log.Printf("err: %w", err)
			break
		}
		return m, domain.UpdateCmd(domain.UpdateMsg{})
	case domain.ChangeMsg:
		err := m.todoRepo.Update(context.TODO(), domain.Item(msg).ID, domain.Item(msg))
		if err != nil {
			// TODO: popup
			log.Printf("err: %w", err)
			break
		}
		return m, domain.UpdateCmd(domain.UpdateMsg{})
	default:
		for i := range m.models {
			cmds = append(cmds, m.update(domain.View(i), msg))
		}
	}

	return m, tea.Batch(cmds...)
}

func (m *mainModel) View() string {
	var s string
	renders := []string{}

	log.Printf("focused view: %v", m.focusedView)

	switch m.focusedView {
	case int(domain.ViewTodoList), int(domain.ViewInProgressList), int(domain.ViewDoneList), int(domain.ViewCollectionSelector):
		if m.focusedView == int(domain.ViewCollectionSelector) {
			renders = append(renders, focusedViewStyle.Render(m.models[domain.ViewCollectionSelector].View()))
		} else {
			renders = append(renders, docStyle.Render(m.models[domain.ViewCollectionSelector].View()))
		}

		if m.focusedView == int(domain.ViewTodoList) {
			renders = append(renders, focusedViewStyle.Render(m.models[domain.ViewTodoList].View()))
		} else {
			renders = append(renders, docStyle.Render(m.models[domain.ViewTodoList].View()))
		}

		if m.focusedView == int(domain.ViewInProgressList) {
			renders = append(renders, focusedViewStyle.Render(m.models[domain.ViewInProgressList].View()))
		} else {
			renders = append(renders, docStyle.Render(m.models[domain.ViewInProgressList].View()))
		}

		if m.focusedView == int(domain.ViewDoneList) {
			renders = append(renders, focusedViewStyle.Render(m.models[domain.ViewDoneList].View()))
		} else {
			renders = append(renders, docStyle.Render(m.models[domain.ViewDoneList].View()))
		}
	case int(domain.ViewItemForm):
		renders = append(renders, docStyle.Render(m.models[domain.ViewItemForm].View()))
	}

	s += lipgloss.JoinHorizontal(lipgloss.Top, renders...)
	return s
}

func (m *mainModel) changeFocusedView() {
	switch m.focusedView {
	case int(domain.ViewCollectionSelector):
		m.focusedView = int(domain.ViewTodoList)
	case int(domain.ViewTodoList):
		m.focusedView = int(domain.ViewInProgressList)
	case int(domain.ViewInProgressList):
		m.focusedView = int(domain.ViewDoneList)
	case int(domain.ViewDoneList):
		m.focusedView = int(domain.ViewCollectionSelector)
	}
}

func (m *mainModel) update(v domain.View, msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd
	m.models[v], cmd = m.models[v].Update(msg)
	return cmd
}
