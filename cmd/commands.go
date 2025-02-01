package main

import (
	tea "github.com/charmbracelet/bubbletea"
	collectionselector "github.com/somatom98/todoist/models/collection_selector"
	todolist "github.com/somatom98/todoist/models/todo_list"
)

type initMsg struct {
	Models   []tea.Model
	InitCmds []tea.Cmd
}

func initCommand() tea.Msg {
	return initMsg{
		Models: []tea.Model{
			collectionselector.New(),
			todolist.New(),
		},
		InitCmds: []tea.Cmd{
			collectionselector.InitCmd,
			todolist.InitCmd,
		},
	}
}
