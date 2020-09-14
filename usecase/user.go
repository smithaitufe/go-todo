package usecase

import (
	"github.com/smithaitufe/go-todo"
)

type UserUsecase interface {
	GetUsers() ([]todo.User, error)
	GetUser(id uint) (*todo.User, error)
	CreateUser(u todo.User) (*todo.User, error)
	UpdateUser(id uint, u *todo.User) (*todo.User, error)
	DeleteUser(id uint) error
}

type userUsecase struct {
	userRepo todo.UserRepository
}
