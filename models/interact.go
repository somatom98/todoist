package models

import (
	"errors"
	"log"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/somatom98/todoist/domain"
)

const (
	title = iota
	description
	collection
)

const (
	hotPink  = lipgloss.Color("#FF06B7")
	darkGray = lipgloss.Color("#767676")
	darkRed  = lipgloss.Color("#CC0000")
)

var (
	errInvalidTitle       = errors.New("invalid title")
	errInvalidDescription = errors.New("invalid description")
	errInvalidCollection  = errors.New("invalid collection")
)

var (
	inputStyle    = lipgloss.NewStyle().Foreground(hotPink)
	continueStyle = lipgloss.NewStyle().Foreground(darkGray)
	errorStyle    = lipgloss.NewStyle().Foreground(darkRed)
)

type interactModel struct {
	inputs       []textinput.Model
	focusedInput int
	item         domain.Item
	operation    domain.Operation
	err          error
	focused      bool
}

func NewInteractModel() *interactModel {
	inputs := make([]textinput.Model, 3)
	inputs[title] = textinput.New()
	inputs[title].Placeholder = "Title - It must be unique"
	inputs[title].Focus()

	inputs[description] = textinput.New()
	inputs[description].Placeholder = "Description - here you can add more details"

	inputs[collection] = textinput.New()
	inputs[collection].Placeholder = "Collection"

	return &interactModel{
		inputs:       inputs,
		focusedInput: 0,
	}
}

func (m *interactModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m *interactModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := make([]tea.Cmd, len(m.inputs))

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if m.focusedInput != len(m.inputs)-1 {
				m.nextInput()
				return m, nil
			}

			item := m.item
			item.Tit = m.inputs[title].Value()
			item.Descr = m.inputs[description].Value()
			item.Collection = domain.Collection(m.inputs[collection].Value())
			m.err = m.validate(item)
			if m.err != nil {
				log.Printf("error: %v", m.err)
				return m, nil
			}

			switch m.operation {
			case domain.OperationAdd:
				cmds = append(cmds, domain.AddCmd(item))
			case domain.OperationChange:
				cmds = append(cmds, domain.ChangeCmd(item))
			}
			return m, tea.Batch(cmds...)
		case tea.KeyShiftTab:
			m.prevInput()
		case tea.KeyTab:
			m.nextInput()
		case tea.KeyEsc:
			// TODO: domain.CancelCmd()
		case tea.KeyCtrlC:
			return m, tea.Quit
		}

		for i := range m.inputs {
			m.inputs[i].Blur()
		}
		m.inputs[m.focusedInput].Focus()
	case domain.OperationMsg:
		m.operation = msg.Operation
		m.item = msg.Item
		m.inputs[collection].SetValue(string(msg.Item.Collection))
		if msg.Operation == domain.OperationChange {
			m.inputs[title].SetValue(msg.Item.Tit)
			m.inputs[description].SetValue(msg.Item.Descr)
		}
	}

	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}
	return m, tea.Batch(cmds...)
}

func (m *interactModel) View() string {
	style := paneStyle
	if m.focused {
		style = focusedPaneStyle
	}

	if m.err != nil {
		style.Render(errorStyle.Render(m.err.Error()))
	}

	return style.Render(m.inputs[m.focusedInput].View())
}

func (m *interactModel) ChangeFocus() {
	m.focused = !m.focused
}

func (m *interactModel) nextInput() {
	m.focusedInput = (m.focusedInput + 1) % len(m.inputs)
}

func (m *interactModel) prevInput() {
	m.focusedInput--
	if m.focusedInput < 0 {
		m.focusedInput = len(m.inputs) - 1
	}
}

func (m interactModel) validate(item domain.Item) error {
	if item.Tit == "" {
		return errInvalidTitle
	}
	if item.Descr == "" {
		return errInvalidDescription
	}
	if item.Collection == "" {
		return errInvalidCollection
	}
	return nil
}
