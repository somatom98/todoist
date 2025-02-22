package domain

type Pane int

const (
	PaneCollectionSelector Pane = iota
	PaneTodoList
	PaneInProgressList
	PaneDoneList
	PaneItemForm
)
