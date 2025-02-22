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

func (c *paneSelector) CurrentFocus() domain.Pane {
	return c.focusedPane
}

func (c *paneSelector) FocusNext() {
	c.focusedPane = orderedPanes[(c.focusedPaneIndex()+1)%len(orderedPanes)]
}

func (c *paneSelector) FocusPrev() {
	prev := c.focusedPaneIndex() - 1
	if prev < 0 {
		prev = len(orderedPanes) - 1
	}
	c.focusedPane = orderedPanes[prev]
}

func (c *paneSelector) focusedPaneIndex() int {
	for i, p := range orderedPanes {
		if c.focusedPane == p {
			return i
		}
	}
	return 0
}

func (c *paneSelector) SetView(view domain.Pane) {
	c.focusedPane = view
}
