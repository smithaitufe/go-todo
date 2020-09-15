package resolver

import (
	"context"
	"fmt"

	"github.com/graph-gophers/graphql-go"
	"github.com/smithaitufe/go-todo"
	"github.com/smithaitufe/go-todo/usecase"
)

type UserInput struct {
	FirstName string
	LastName  string
	Email     string
}

type userResolver struct {
	user *todo.User
}

func (r userResolver) ID() graphql.ID {
	return graphql.ID(fmt.Sprint(r.user.ID))
}
func (r userResolver) FirstName() string {
	return r.user.FirstName
}
func (r userResolver) LastName() string {
	return r.user.LastName
}
func (r userResolver) Email() string {
	return r.user.Email
}
func userDto(in UserInput) todo.User {
	return todo.User{
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Email:     in.Email,
	}
}
func createUser(ctx context.Context, u usecase.UserUsecase, in UserInput) (*userResolver, error) {
	o, err := u.CreateUser(userDto(in))
	if err != nil {
		return nil, err
	}
	return &userResolver{o}, nil
}
func (s *rootResolver) CreateUser(ctx context.Context, args struct {
	Input UserInput
}) (*userResolver, error) {
	return createUser(ctx, s.userUsecase, args.Input)
}

func users(ctx context.Context, u usecase.UserUsecase) ([]userResolver, error) {
	o, err := u.GetUsers()
	if err != nil {
		return nil, err
	}

	r := make([]userResolver, len(o))
	for i := range o {
		r[i] = userResolver{&o[i]}
	}
	return r, nil
}
func (s *rootResolver) Users(ctx context.Context) ([]userResolver, error) {
	return users(ctx, s.userUsecase)
}

func user(ctx context.Context, u usecase.UserUsecase, id int32) (*userResolver, error) {
	o, err := u.GetUser(id)
	if err != nil {
		return nil, err
	}
	return &userResolver{o}, nil
}
func (s *rootResolver) User(ctx context.Context, args struct {
	ID string
}) (*userResolver, error) {
	return user(ctx, s.userUsecase, strint32(args.ID))
}
