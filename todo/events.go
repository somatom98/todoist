package todo

import (
	tea "github.com/charmbracelet/bubbletea"
)

type UpdateMsg struct {
	Collection *Collection
}

var _ tea.Cmd = UpdateCmd(UpdateMsg{})

func UpdateCmd(msg UpdateMsg) func() tea.Msg {
	return func() tea.Msg { return msg }
}
