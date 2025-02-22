package controllers

import "github.com/somatom98/todoist/domain"

var orderedPanes = []domain.Pane{
	domain.PaneCollectionSelector,
	domain.PaneTodoList,
	domain.PaneInProgressList,
	domain.PaneDoneList,
}

type paneSelector struct {
	focusedPane domain.Pane
}

func NewPaneSelector() *paneSelector {
	return &paneSelector{
		focusedPane: orderedPanes[0],
	}
}

func (c *paneSelector) GetFocus() domain.Pane {
	return c.focusedPane
}

func (c *paneSelector) FocusNext() {
	current := 0
	for i, p := range orderedPanes {
		if c.focusedPane == p {
			current = i
			break
		}
	}

	c.focusedPane = orderedPanes[(current+1)%len(orderedPanes)]
}

func (c *paneSelector) SetView(view domain.Pane) {
	c.focusedPane = view
}
