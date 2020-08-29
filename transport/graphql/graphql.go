package graphql

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/smithaitufe/go-todo"
	"github.com/smithaitufe/go-todo/transport/graphql/schema"
)

type graphqlServer struct {
	todoUsecase todo.TodoUsecase
}

func NewGraphQLServer(todoUsecase todo.TodoUsecase) *relay.Handler {
	s := schema.GetRootSchema()
	opts := []graphql.SchemaOpt{graphql.UseFieldResolvers()}
	parsedSchema := graphql.MustParseSchema(s, &graphqlServer{todoUsecase}, opts...)
	return &relay.Handler{Schema: parsedSchema}

}
