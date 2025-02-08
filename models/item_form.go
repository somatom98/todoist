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
	inputs    []textinput.Model
	focused   int
	item      todo.Item
	operation todo.Operation
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
			if m.focused != len(m.inputs)-1 {
				m.nextInput()
			}

			item := m.item
			item.Tit = m.inputs[title].Value()
			item.Descr = m.inputs[description].Value()
			item.Collection = todo.Collection(m.inputs[collection].Value())

			cmds = append(cmds, ViewCmd(ViewMsg{View: viewCollectionSelector}))
			switch m.operation {
			case todo.OperationAdd:
				cmds = append(cmds, todo.AddCmd(item))
			case todo.OperationChange:
				cmds = append(cmds, todo.ChangeCmd(item))
			}
			return m, tea.Batch(cmds...)
		case tea.KeyShiftTab:
			m.prevInput()
		case tea.KeyTab:
			m.nextInput()
		case tea.KeyEsc:
			cmds = append(cmds, ViewCmd(ViewMsg{View: viewCollectionSelector}))
			return m, tea.Batch(cmds...)
		default:
			log.Printf("collection: %s", m.inputs[collection].Value())
		}

		for i := range m.inputs {
			m.inputs[i].Blur()
		}
		m.inputs[m.focused].Focus()
	case todo.OperationMsg:
		m.operation = msg.Operation
		m.item = msg.Item
		m.inputs[collection].SetValue(string(msg.Item.Collection))
		if msg.Operation == todo.OperationChange {
			m.inputs[title].SetValue(msg.Item.Tit)
			m.inputs[description].SetValue(msg.Item.Descr)
		}
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
		inputStyle.Render("Description"),
		m.inputs[description].View(),
		inputStyle.Render("Collection"),
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
