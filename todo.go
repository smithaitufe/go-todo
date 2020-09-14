package todo

import (
	"fmt"

	"github.com/graph-gophers/graphql-go"
)

type Todo struct {
	ID          uint   `gorm:"primary_key"`
	UserID      uint   `gorm:"not null"`
	User        User   `gorm:"foreignkey:UserID;association_foreignkey:ID"`
	Description string `gorm:"not null;"`
	Completed   bool   `gorm:"default:false"`
}

func (t Todo) Id() graphql.ID {
	return graphql.ID(fmt.Sprintf("%d", t.ID))
}

type TodoRepository interface {
	GetTodos() ([]Todo, error)
	GetTodo(id int64) (*Todo, error)
	CreateTodo(t Todo) (*Todo, error)
	UpdateTodo(id int64, t Todo) (*Todo, error)
	DeleteTodo(t Todo) (*Todo, error)
}
