package todoist

import tea "github.com/charmbracelet/bubbletea"

type InitMsg struct{}

func InitCmd() tea.Msg {
	return InitMsg{}
}
