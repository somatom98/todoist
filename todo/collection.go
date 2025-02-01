package todo

import (
	"github.com/charmbracelet/bubbles/list"
)

type Collection string

var _ list.Item = Collection("")

// Implement list.Item interface

func (c Collection) FilterValue() string {
	return string(c)
}

func (c Collection) Title() string {
	return string(c)
}

func (c Collection) Description() string {
	return ""
}
