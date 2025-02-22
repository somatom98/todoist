package todo

type Status string

const (
	StatusTodo       = "todo"
	StatusInProgress = "in_progress"
	StatusDone       = "done"
)

var statuses = []Status{
	StatusTodo,
	StatusInProgress,
	StatusDone,
}

func (status Status) Next() Status {
	for i, s := range statuses {
		if s == status && i < len(statuses)-1 {
			return statuses[i+1]
		}
	}
	return StatusDone
}
