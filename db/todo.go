package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/smithaitufe/go-todo"
)

type TodoRepository interface {
	GetTodos() ([]todo.Todo, error)
	GetTodo(id int64) (*todo.Todo, error)
	CreateTodo(t todo.Todo) (*todo.Todo, error)
	UpdateTodo(id int64, t todo.Todo) (*todo.Todo, error)
	DeleteTodo(t todo.Todo) (*todo.Todo, error)
}

type todoRepository struct {
	DB *gorm.DB
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return todoRepository{db}
}
func (s todoRepository) GetTodos() ([]todo.Todo, error) {
	var todos []todo.Todo
	err := s.DB.Find(&todos).Error
	if err != nil {
		return nil, err
	}
	fmt.Printf("%#v", todos)

	return todos, nil
}

func (s todoRepository) GetTodo(id int64) (*todo.Todo, error) {
	var todo todo.Todo
	err := s.DB.Where(map[string]interface{}{"id": id}).Find(&todo).Error
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func (s todoRepository) CreateTodo(t todo.Todo) (*todo.Todo, error) {
	err := s.DB.Create(&t).Error
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (s todoRepository) UpdateTodo(id int64, t todo.Todo) (*todo.Todo, error) {
	err := s.DB.Model(&todo.Todo{}).Where(map[string]interface{}{"id": id}).Updates(&t).Error
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (s todoRepository) DeleteTodo(todo todo.Todo) (*todo.Todo, error) {
	err := s.DB.Delete(&todo).Error
	if err != nil {
		return nil, err
	}
	return &todo, nil
}
