package models

import (
	tea "github.com/charmbracelet/bubbletea"
)

type ViewMsg struct {
	View view
}

var _ tea.Cmd = ViewCmd(ViewMsg{})

func ViewCmd(msg ViewMsg) func() tea.Msg {
	return func() tea.Msg { return msg }
}
