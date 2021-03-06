package todo

import (
	"fmt"

	"github.com/graph-gophers/graphql-go"
)

type Todo struct {
	ID          uint   `gorm:"primary_key"`
	Description string `gorm:"not null;"`
	Completed   bool   `gorm:"default:false"`
}

func (t Todo) Id() graphql.ID {
	return graphql.ID(fmt.Sprintf("%d", t.ID))
}

type TodoService interface {
	FindTodos() ([]Todo, error)
	FindTodo(id int64) (*Todo, error)
	CreateTodo(t Todo) (*Todo, error)
	UpdateTodo(id int64, t Todo) (*Todo, error)
	DeleteTodo(t Todo) (*Todo, error)
}
