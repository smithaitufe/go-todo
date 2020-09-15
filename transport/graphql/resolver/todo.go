package resolver

import (
	"context"
	"fmt"

	"github.com/graph-gophers/graphql-go"
	"github.com/smithaitufe/go-todo"
)

type todoResolver struct {
	todo *todo.Todo
}

func (r todoResolver) ID() graphql.ID {
	return graphql.ID(fmt.Sprintf("%d", r.todo.ID))
}
func (r todoResolver) UserID() graphql.ID {
	return graphql.ID(fmt.Sprintf("%d", r.todo.UserID))
}
func (r todoResolver) Description() string {
	return r.todo.Description
}
func (r todoResolver) Completed() bool {
	return r.todo.Completed
}
func (r todoResolver) CreatedAt() graphql.Time {
	return graphql.Time{Time: r.todo.CreatedAt}
}
func (r todoResolver) UpdatedAt() graphql.Time {
	return graphql.Time{Time: r.todo.UpdatedAt}
}
func (s *rootResolver) Todos() ([]*todoResolver, error) {
	todos, err := s.todoUsecase.FetchTodos()
	if err != nil {
		return nil, err
	}
	r := make([]*todoResolver, len(todos))
	for i := range todos {
		r[i] = &todoResolver{&todos[i]}
	}
	return r, nil
}

func (s *rootResolver) Todo(ctx context.Context, args struct {
	ID string
}) (*todoResolver, error) {
	todo, err := s.todoUsecase.FetchTodo(strint32(args.ID))
	if err != nil {
		return nil, err
	}
	return &todoResolver{todo}, nil
}

type TodoInput struct {
	UserID      string
	Description string
	Completed   bool
}

func todoDto(in TodoInput) todo.Todo {
	return todo.Todo{
		UserID:      strint32(in.UserID),
		Description: in.Description,
		Completed:   in.Completed,
	}
}

func (s *rootResolver) AddTodo(ctx context.Context, args struct {
	Input TodoInput
}) (*todoResolver, error) {
	t, err := s.todoUsecase.AddTodo(todoDto(args.Input))
	if err != nil {
		return nil, err

	}
	return &todoResolver{t}, nil
}
func (s *rootResolver) UpdateTodo(ctx context.Context, args struct {
	ID    string
	Input TodoInput
}) (*todoResolver, error) {

	t, err := s.todoUsecase.UpdateTodo(strint32(args.ID), todoDto(args.Input))
	if err != nil {
		return nil, err

	}
	return &todoResolver{t}, nil
}
func (s *rootResolver) DeleteTodo(ctx context.Context, args struct {
	ID string
}) (*todoResolver, error) {
	t, err := s.todoUsecase.DeleteTodo(strint32(args.ID))
	if err != nil {
		return nil, err

	}
	return &todoResolver{t}, nil
}
