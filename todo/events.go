package todo

import (
	tea "github.com/charmbracelet/bubbletea"
)

type UpdateMsg struct{}

func UpdateCmd() tea.Msg {
	return UpdateMsg{}
}
