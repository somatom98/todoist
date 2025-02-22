package domain

type View int

const (
	ViewCollectionSelector View = iota
	ViewTodoList
	ViewInProgressList
	ViewDoneList
	ViewItemForm
)

var Views = []View{}
