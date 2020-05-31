package main

import (
	"log"
	"net/http"

	"github.com/smithaitufe/go-todo/db"
	"github.com/smithaitufe/go-todo/transport/graphql"
	"github.com/smithaitufe/go-todo/usecase"
)

func main() {
	connection := db.Connect(db.Config{
		Host:     "todo-db",
		Port:     4321,
		Database: "postgres",
		Password: "postgres",
		User:     "postgres",
	})

	todoRepository := db.NewTodoRepository(connection)
	todoUsecase := usecase.NewTodoUsecase(todoRepository)

	graphqlServer := graphql.NewGraphQLServer(todoUsecase)
	http.Handle("/graphql", graphqlServer)
	log.Fatal(http.ListenAndServe(":5000", nil))
}
