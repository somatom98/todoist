package models

type view int

const (
	viewCollectionSelector view = iota
	viewTodoList
	viewInProgressList
	viewDoneList
	viewItemForm
)
