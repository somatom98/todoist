package models

import (
	"errors"
	"fmt"
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

type itemFormModel struct {
	inputs    []textinput.Model
	focused   int
	item      domain.Item
	operation domain.Operation
	err       error
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
			item.Collection = domain.Collection(m.inputs[collection].Value())
			m.err = m.validate(item)
			log.Printf("error: %v", m.err)
			if m.err != nil {
				break
			}

			cmds = append(cmds, ViewCmd(ViewMsg{View: domain.ViewCollectionSelector}))
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
			cmds = append(cmds, ViewCmd(ViewMsg{View: domain.ViewCollectionSelector}))
			return m, tea.Batch(cmds...)
		}

		for i := range m.inputs {
			m.inputs[i].Blur()
		}
		m.inputs[m.focused].Focus()
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

func (m itemFormModel) View() string {
	bottomMessage := continueStyle.Render("Continue ->")
	if m.err != nil {
		bottomMessage = errorStyle.Render(fmt.Sprintf("Error: %s", m.err))
	}

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
		bottomMessage,
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

func (m itemFormModel) validate(item domain.Item) error {
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
