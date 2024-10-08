package cmd

// TodoEntity represents a task with a description, due date, and status.
type TodoEntity struct {
	Description string
	DueDate     string
	IsDone      bool
}

// NewTodoEntity is the constructor for TodoEntity with default values.
func NewTodoEntity(description string, dueDate string, isDone bool) *TodoEntity {
	return &TodoEntity{
		Description: description,
		DueDate:     dueDate,
		IsDone:      isDone,
	}
}

// SetDone changes the status of IsDone to true.
func (t *TodoEntity) SetDone() {
	t.IsDone = true
}
