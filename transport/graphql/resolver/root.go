package resolver

import (
	"strconv"

	"github.com/smithaitufe/go-todo/usecase"
)

type rootResolver struct {
	todoUsecase usecase.TodoUsecase
	userUsecase usecase.UserUsecase
}

func NewRootResolver(t usecase.TodoUsecase, u usecase.UserUsecase) rootResolver {
	return rootResolver{
		todoUsecase: t,
		userUsecase: u,
	}
}
func strint32(id string) int32 {
	i, _ := strconv.Atoi(id)
	return int32(i)
}
