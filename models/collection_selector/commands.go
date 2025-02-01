package collectionselector

import (
	tea "github.com/charmbracelet/bubbletea"
)

func InitCmd() tea.Msg {
	return getCollectionsCmd()
}

type getCollectionsMsg struct{}

func getCollectionsCmd() tea.Msg {
	return getCollectionsMsg{}
}
