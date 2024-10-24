package cmd

// TodoEntity represents a task with a description, due date, and status.
type TodoEntity struct {
	Name    string `json:"name"`
	DueDate string `json:"dueDate"`
	IsDone  bool   `json:"isDone"`
}

// NewTodoEntity is the constructor for TodoEntity with default values.
func NewTodoEntity(name string, dueDate string, isDone bool) *TodoEntity {
	return &TodoEntity{
		Name:    name,
		DueDate: dueDate,
		IsDone:  isDone,
	}
}

// SetDone changes the status of IsDone to true.
func (t *TodoEntity) SetDone() {
	t.IsDone = true
}
