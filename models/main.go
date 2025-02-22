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
	docStyle = lipgloss.NewStyle().
			Margin(2, 2)
	paneStyle = docStyle.
			Border(lipgloss.NormalBorder()).
			Padding(1, 2)
	focusedPaneStyle = docStyle.
				Border(lipgloss.DoubleBorder()).
				BorderForeground(lipgloss.AdaptiveColor{Light: "#EE6FF8", Dark: "#EE6FF8"}).
				Padding(1, 2)
)

type mainModel struct {
	todoRepo     controllers.ItemsRepo
	paneSelector controllers.PaneSelector
	models       []tea.Model
	collection   domain.Collection
}

func NewMain(todoRepo controllers.ItemsRepo, paneSelector controllers.PaneSelector) *mainModel {
	return &mainModel{
		todoRepo:     todoRepo,
		paneSelector: paneSelector,
		models: []tea.Model{
			domain.PaneCollectionSelector: NewCollectionSelector(todoRepo),
			domain.PaneTodoList:           NewTaskList(domain.StatusTodo, todoRepo),
			domain.PaneInProgressList:     NewTaskList(domain.StatusInProgress, todoRepo),
			domain.PaneDoneList:           NewTaskList(domain.StatusDone, todoRepo),
			domain.PaneItemForm:           NewInteractModel(),
		},
	}
}

func (m *mainModel) Init() tea.Cmd {
	m.models[m.paneSelector.CurrentFocus()].(Model).ChangeFocus()
	return domain.UpdateCmd(domain.UpdateMsg{})
}

func (m *mainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := []tea.Cmd{}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.paneSelector.CurrentFocus() == domain.PaneItemForm {
			cmds = append(cmds, m.update(domain.PaneItemForm, msg))
			break
		}

		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "tab", "l":
			m.changeFocus(m.paneSelector.FocusNext)
		case "h":
			m.changeFocus(m.paneSelector.FocusPrev)
		default:
			cmds = append(cmds, m.update(m.paneSelector.CurrentFocus(), msg))
		}
	case domain.OperationMsg:
		m.setFocus(domain.PaneItemForm)
	case domain.AddMsg:
		err := m.todoRepo.Add(context.TODO(), domain.Item(msg))
		if err != nil {
			log.Printf("err: %v", err)
			break
		}
		m.setFocus(domain.PaneCollectionSelector)
		return m, domain.UpdateCmd(domain.UpdateMsg{})
	case domain.ChangeMsg:
		err := m.todoRepo.Update(context.TODO(), domain.Item(msg).ID, domain.Item(msg))
		if err != nil {
			log.Printf("err: %v", err)
			break
		}
		m.setFocus(domain.PaneCollectionSelector)
		return m, domain.UpdateCmd(domain.UpdateMsg{})
	default:
		for i := range m.models {
			cmds = append(cmds, m.update(domain.Pane(i), msg))
		}
	}

	return m, tea.Batch(cmds...)
}

func (m *mainModel) View() string {
	return lipgloss.JoinHorizontal(lipgloss.Left,
		m.models[domain.PaneCollectionSelector].View(),
		lipgloss.JoinVertical(lipgloss.Top,
			m.models[domain.PaneItemForm].View(),
			lipgloss.JoinHorizontal(lipgloss.Left,
				m.models[domain.PaneTodoList].View(),
				m.models[domain.PaneInProgressList].View(),
				m.models[domain.PaneDoneList].View(),
			),
		),
	)
}

func (m *mainModel) update(v domain.Pane, msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd
	m.models[v], cmd = m.models[v].Update(msg)
	return cmd
}

func (m *mainModel) changeFocus(action func()) {
	m.models[m.paneSelector.CurrentFocus()].(Model).ChangeFocus()
	action()
	m.models[m.paneSelector.CurrentFocus()].(Model).ChangeFocus()
}

func (m *mainModel) setFocus(pane domain.Pane) {
	m.models[m.paneSelector.CurrentFocus()].(Model).ChangeFocus()
	m.paneSelector.SetFocus(pane)
	m.models[m.paneSelector.CurrentFocus()].(Model).ChangeFocus()
}
