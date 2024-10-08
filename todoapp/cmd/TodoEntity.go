package cmd

// TodoEntity represents a task with a description, due date, and status.
type TodoEntity struct {
	Description string
	DueDate     string
	IsDone      bool
}

// NewTodoEntity is the constructor for TodoEntity with default values.
func NewTodoEntity() *TodoEntity {
	return &TodoEntity{
		Description: "",
		DueDate:     "",
		IsDone:      false,
	}
}

// SetDone changes the status of IsDone to true.
func (t *TodoEntity) SetDone(isDone bool) {
	t.IsDone = isDone
}
