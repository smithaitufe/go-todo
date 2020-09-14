package graphql

import (
	"context"
	"strconv"

	"github.com/graph-gophers/graphql-go"
	"github.com/smithaitufe/go-todo"
)

// type TodoResolver struct {
// 	t *todo.Todo
// }

// func (r *TodoResolver) ID() graphql.ID {
// 	return graphql.ID(fmt.Sprintf("%d", r.t.ID))
// }
// func (r *TodoResolver) Description() string {
// 	return r.t.Description
// }
// func (r *TodoResolver) Completed() bool {
// 	return r.t.Completed
// }
func (s *graphqlServer) Todos() ([]*todo.Todo, error) {
	todos, err := s.todoUsecase.FetchTodos()
	if err != nil {
		return nil, err
	}
	out := make([]*todo.Todo, len(todos))
	for i := range todos {
		out[i] = &todos[i]
	}

	return out, nil
}

func (s *graphqlServer) Todo(ctx context.Context, args struct {
	ID graphql.ID
}) (*todo.Todo, error) {

	todo, err := s.todoUsecase.FetchTodo(int64ID(args.ID))
	if err != nil {
		return nil, err
	}
	return todo, nil
}

type AddTodoInput struct {
	Description string
}

func (s *graphqlServer) AddTodo(ctx context.Context, args struct {
	Input AddTodoInput
}) (*todo.Todo, error) {
	td := todo.Todo{
		Description: args.Input.Description,
	}
	t, err := s.todoUsecase.AddTodo(td)
	if err != nil {
		return nil, err

	}
	return t, nil
}

type UpdateTodoInput struct {
	Description string
	Completed   bool
}

func (s *graphqlServer) UpdateTodo(ctx context.Context, args struct {
	ID    graphql.ID
	Input UpdateTodoInput
}) (*todo.Todo, error) {
	td := todo.Todo{
		Description: args.Input.Description,
		Completed:   args.Input.Completed,
	}
	t, err := s.todoUsecase.UpdateTodo(int64ID(args.ID), td)
	if err != nil {
		return nil, err

	}
	return t, nil
}
func (s *graphqlServer) DeleteTodo(ctx context.Context, args struct {
	ID graphql.ID
}) (*todo.Todo, error) {
	t, err := s.todoUsecase.DeleteTodo(int64ID(args.ID))
	if err != nil {
		return nil, err

	}
	return t, nil
}
func int64ID(id graphql.ID) int64 {
	i, _ := strconv.ParseInt(string(id), 10, 64)
	return i
}
