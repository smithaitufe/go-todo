package usecase

import (
	"github.com/smithaitufe/go-todo"
)

type UserUsecase interface {
	GetUsers() ([]todo.User, error)
	GetUser(id uint) (*todo.User, error)
	CreateUser(u todo.User) (*todo.User, error)
	UpdateUser(id uint, u todo.User) (*todo.User, error)
	DeleteUser(id uint) error
}

type userUsecase struct {
	userRepo todo.UserRepository
}

func NewUserUsecase(ur todo.UserRepository) UserUsecase {
	return &userUsecase{ur}
}

func (u *userUsecase) GetUsers() ([]todo.User, error) {
	o, err := u.userRepo.FindUsers()
	if err != nil {
		return nil, err
	}
	return o, nil
}
func (u *userUsecase) GetUser(id uint) (*todo.User, error) {
	return u.userRepo.FindUser(id)
}
func (u *userUsecase) CreateUser(in todo.User) (*todo.User, error) {
	return u.userRepo.CreateUser(in)
}

func (u *userUsecase) UpdateUser(id uint, in todo.User) (*todo.User, error) {
	o, err := u.userRepo.FindUser(id)
	if err != nil {
		return nil, err
	}
	in.ID = o.ID
	return u.userRepo.UpdateUser(in)
}

func (u *userUsecase) DeleteUser(id uint) error {
	o, err := u.userRepo.FindUser(id)
	if err != nil {
		return err
	}
	return u.userRepo.DeleteUser(*o)
}
