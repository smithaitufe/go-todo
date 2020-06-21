package usecase

import (
	"github.com/smithaitufe/go-todo"

	"github.com/smithaitufe/go-todo/db"
)

type todoUsecase struct {
	todoRepository db.TodoRepository
}

func NewTodoUsecase(todoRepository db.TodoRepository) todo.TodoUsecase {
	return &todoUsecase{todoRepository}
}
func (u *todoUsecase) FetchTodos() ([]todo.Todo, error) {
	todos, err := u.todoRepository.GetTodos()
	if err != nil {
		return nil, err
	}
	return todos, nil
}
func (u *todoUsecase) FetchTodo(id int64) (*todo.Todo, error) {
	todo, err := u.todoRepository.GetTodo(id)
	if err != nil {
		return nil, err
	}
	return todo, nil
}
func (u *todoUsecase) AddTodo(t todo.Todo) (*todo.Todo, error) {
	todo, err := u.todoRepository.CreateTodo(t)
	if err != nil {
		return nil, err
	}
	return todo, nil
}
func (u *todoUsecase) UpdateTodo(id int64, t todo.Todo) (*todo.Todo, error) {
	todo, err := u.todoRepository.UpdateTodo(id, t)
	if err != nil {
		return nil, err
	}
	return todo, nil
}
func (u *todoUsecase) DeleteTodo(id int64) (*todo.Todo, error) {
	todo, err := u.todoRepository.GetTodo(id)
	if err != nil {
		return nil, err
	}
	todo, err = u.todoRepository.DeleteTodo(*todo)
	if err != nil {
		return nil, err
	}
	return todo, nil
}
