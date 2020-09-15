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
		Host:     "localhost",
		Port:     5432,
		Database: "go_todo",
		Password: "postgres",
		User:     "postgres",
	}).Debug()

	defer connection.Close()

	tr := db.NewTodoRepository(connection)
	tu := usecase.NewTodoUsecase(tr)

	ur := db.NewUserRepository(connection)
	uu := usecase.NewUserUsecase(ur)

	graphqlServer := graphql.NewGraphQLServer(tu, uu)
	http.Handle("/graphql", graphqlServer)
	log.Fatal(http.ListenAndServe(":5000", nil))
}
