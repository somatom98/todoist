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
			domain.PaneItemForm:           NewItemFormModel(),
		},
	}
}

func (m *mainModel) Init() tea.Cmd {
	m.models[m.paneSelector.CurrentFocus()].(Model).ChangeFocus()
	return domain.UpdateCmd(domain.UpdateMsg{})
}

func (m *mainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	log.Printf("MAIN, msg: %T - %+v", msg, msg)
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
			m.models[m.paneSelector.CurrentFocus()].(Model).ChangeFocus()
			m.paneSelector.FocusNext()
			m.models[m.paneSelector.CurrentFocus()].(Model).ChangeFocus()
		case "h":
			m.models[m.paneSelector.CurrentFocus()].(Model).ChangeFocus()
			m.paneSelector.FocusPrev()
			m.models[m.paneSelector.CurrentFocus()].(Model).ChangeFocus()
		default:
			cmds = append(cmds, m.update(m.paneSelector.CurrentFocus(), msg))
		}
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
			cmds = append(cmds, m.update(domain.Pane(i), msg))
		}
	}

	return m, tea.Batch(cmds...)
}

func (m *mainModel) View() string {
	var s string
	renders := []string{}

	log.Printf("focused view: %v", m.paneSelector.CurrentFocus())

	switch m.paneSelector.CurrentFocus() {
	case domain.PaneTodoList, domain.PaneInProgressList, domain.PaneDoneList, domain.PaneCollectionSelector:
		renders = append(renders, m.models[domain.PaneCollectionSelector].View())
		renders = append(renders, m.models[domain.PaneTodoList].View())
		renders = append(renders, m.models[domain.PaneInProgressList].View())
		renders = append(renders, m.models[domain.PaneDoneList].View())
	case domain.PaneItemForm:
		renders = append(renders, docStyle.Render(m.models[domain.PaneItemForm].View()))
	}

	s = lipgloss.JoinHorizontal(lipgloss.Left, renders...)
	return s
}

func (m *mainModel) update(v domain.Pane, msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd
	m.models[v], cmd = m.models[v].Update(msg)
	return cmd
}
