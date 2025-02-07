package models

import (
	"fmt"
	"log"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/somatom98/todoist/todo"
)

const (
	title = iota
	description
	collection
)

const (
	hotPink  = lipgloss.Color("#FF06B7")
	darkGray = lipgloss.Color("#767676")
)

var (
	inputStyle    = lipgloss.NewStyle().Foreground(hotPink)
	continueStyle = lipgloss.NewStyle().Foreground(darkGray)
)

type itemFormModel struct {
	inputs  []textinput.Model
	focused int
	err     error
}

func NewItemFormModel() itemFormModel {
	inputs := make([]textinput.Model, 3)
	inputs[title] = textinput.New()
	inputs[title].Placeholder = "Title - It must be unique"
	inputs[title].Focus()
	inputs[title].Width = 30

	inputs[description] = textinput.New()
	inputs[description].Placeholder = "Description - here you can add more details"

	inputs[collection] = textinput.New()
	inputs[collection].Placeholder = "Collection"

	return itemFormModel{
		inputs:  inputs,
		focused: 0,
		err:     nil,
	}
}

func (m itemFormModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m itemFormModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	log.Printf("ITEM FORM, msg: %T - %+v", msg, msg)
	cmds := make([]tea.Cmd, len(m.inputs))

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if m.focused == len(m.inputs)-1 {
				cmds = append(cmds, ViewCmd(ViewMsg{View: viewTodoList}))
				cmds = append(cmds, todo.AddCmd(todo.AddMsg(
					todo.New(
						m.inputs[title].Value(),
						m.inputs[description].Value(),
						m.inputs[collection].Value()))))
				return m, tea.Batch(cmds...)
			}
			m.nextInput()
		case tea.KeyShiftTab:
			m.prevInput()
		case tea.KeyTab:
			m.nextInput()
		}

		for i := range m.inputs {
			m.inputs[i].Blur()
		}
		m.inputs[m.focused].Focus()
	case todo.UpdateMsg:
		if msg.Collection == nil {
			break
		}
		m.inputs[collection].Placeholder = string(*msg.Collection)
	}

	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}
	return m, tea.Batch(cmds...)
}

func (m itemFormModel) View() string {
	return fmt.Sprintf(
		`
 %s
 %s

 %s
 %s

 %s
 %s

 %s
`,
		inputStyle.Width(30).Render("Title"),
		m.inputs[title].View(),
		inputStyle.Width(6).Render("Description"),
		m.inputs[description].View(),
		inputStyle.Width(6).Render("Collection"),
		m.inputs[collection].View(),
		continueStyle.Render("Continue ->"),
	) + "\n"
}

func (m *itemFormModel) nextInput() {
	m.focused = (m.focused + 1) % len(m.inputs)
}

func (m *itemFormModel) prevInput() {
	m.focused--
	if m.focused < 0 {
		m.focused = len(m.inputs) - 1
	}
}
