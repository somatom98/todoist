package models

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/somatom98/todoist/domain"
)

type ViewMsg struct {
	View domain.View
}

var _ tea.Cmd = ViewCmd(ViewMsg{})

func ViewCmd(msg ViewMsg) func() tea.Msg {
	return func() tea.Msg { return msg }
}
