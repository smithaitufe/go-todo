package todo

type Todo struct {
	ID          int64
	Description string
	Completed   bool
}

type TodoService interface {
	FindTodos() ([]Todo, error)
	FindTodo(id int64) (*Todo, error)
	CreateTodo(t Todo) (*Todo, error)
	UpdateTodo(id int64, t Todo) (*Todo, error)
	DeleteTodo(id int64) (*Todo, error)
}
