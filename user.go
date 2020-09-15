package todo

import (
	"time"
)

type User struct {
	ID        int32  `gorm:"primary_key"`
	FirstName string `gorm:"not nll"`
	LastName  string `gorm:"not nll"`
	Email     string `gorm:"unique;not nll"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Todos     []Todo
}

type UserRepository interface {
	FindUsers() ([]User, error)
	FindUser(id int32) (*User, error)
	CreateUser(u User) (*User, error)
	UpdateUser(u User) (*User, error)
	DeleteUser(u User) error
}
