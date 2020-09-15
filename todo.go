package todo

import (
	"time"
)

type Todo struct {
	ID          int32  `gorm:"primary_key"`
	UserID      int32  `gorm:"not null"`
	User        User   `gorm:"foreignkey:UserID;association_foreignkey:ID"`
	Description string `gorm:"not null;"`
	Completed   bool   `gorm:"default:false"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type TodoRepository interface {
	GetTodos() ([]Todo, error)
	GetTodo(id int32) (*Todo, error)
	CreateTodo(t Todo) (*Todo, error)
	UpdateTodo(id int32, t Todo) (*Todo, error)
	DeleteTodo(t Todo) (*Todo, error)
}
