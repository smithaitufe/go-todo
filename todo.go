package todo

type Todo struct {
	ID          uint64 `gorm:"primary_key;auto_increment"`
	Description string `gorm:"not null;"`
	Completed   bool   `gorm:"default:false"`
}

type TodoService interface {
	FindTodos() ([]Todo, error)
	FindTodo(id int64) (*Todo, error)
	CreateTodo(t Todo) (*Todo, error)
	UpdateTodo(id int64, t Todo) (*Todo, error)
	DeleteTodo(t Todo) (*Todo, error)
}
