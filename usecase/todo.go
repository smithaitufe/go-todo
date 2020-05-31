package usecase

import (
	"github.com/smithaitufe/go-todo"

	"github.com/smithaitufe/go-todo/db"
)

type TodoUsecase struct {
	todoRepository db.TodoRepository
}

func NewTodoUsecase(todoRepository db.TodoRepository) TodoUsecase {
	return TodoUsecase{todoRepository}
}
func (s *TodoUsecase) FetchTodos() ([]todo.Todo, error) {
	todos, err := s.todoRepository.GetTodos()
	if err != nil {
		return nil, err
	}
	return todos, nil
}
func (s *TodoUsecase) FetchTodo(id int64) (*todo.Todo, error) {
	todo, err := s.todoRepository.GetTodo(id)
	if err != nil {
		return nil, err
	}
	return todo, nil
}
func (s *TodoUsecase) AddTodo(t todo.Todo) (*todo.Todo, error) {
	todo, err := s.todoRepository.CreateTodo(t)
	if err != nil {
		return nil, err
	}
	return todo, nil
}
func (s *TodoUsecase) UpdateTodo(id int64, t todo.Todo) (*todo.Todo, error) {
	todo, err := s.todoRepository.UpdateTodo(id, t)
	if err != nil {
		return nil, err
	}
	return todo, nil
}
func (s *TodoUsecase) DeleteTodo(id int64) (*todo.Todo, error) {
	todo, err := s.todoRepository.GetTodo(id)
	if err != nil {
		return nil, err
	}
	todo, err = s.todoRepository.DeleteTodo(*todo)
	if err != nil {
		return nil, err
	}
	return todo, nil
}
