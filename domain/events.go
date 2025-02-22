package domain

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

type AddMsg Item

var _ tea.Cmd = AddCmd(Item{})

func AddCmd(item Item) func() tea.Msg {
	return func() tea.Msg { return AddMsg(item) }
}

type ChangeMsg Item

var _ tea.Cmd = ChangeCmd(Item{})

func ChangeCmd(item Item) func() tea.Msg {
	return func() tea.Msg { return ChangeMsg(item) }
}

type Operation int

const (
	OperationAdd = iota
	OperationChange
)

type OperationMsg struct {
	Operation Operation
	Item      Item
}

var _ tea.Cmd = OperationCmd(OperationAdd, Item{})

func OperationCmd(op Operation, item Item) func() tea.Msg {
	return func() tea.Msg { return OperationMsg{Operation: op, Item: item} }
}
