package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/smithaitufe/go-todo"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) todo.UserRepository {
	return userRepository{db}
}
func (s userRepository) FindUsers() ([]todo.User, error) {
	var u []todo.User
	err := s.db.Model(&todo.User{}).Find(&u).Error
	if err != nil {
		return nil, err
	}
	fmt.Printf("%#v", u)

	return u, nil
}

func (s userRepository) FindUser(id uint) (*todo.User, error) {
	var u todo.User
	err := s.db.Where(map[string]interface{}{"id": id}).Find(&u).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (s userRepository) CreateUser(u todo.User) (*todo.User, error) {
	err := s.db.Model(&todo.User{}).Create(&u).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (s userRepository) UpdateUser(u todo.User) (*todo.User, error) {
	err := s.db.Model(&u).Updates(&u).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (s userRepository) DeleteUser(u todo.User) error {
	return s.db.Model(&u).Delete(&u).Error

}
