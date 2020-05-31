package graphql

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/smithaitufe/go-todo/transport/graphql/schema"
	"github.com/smithaitufe/go-todo/usecase"
)

type graphqlServer struct {
	todoUsecase usecase.TodoUsecase
}

func NewGraphQLServer(todoUsecase usecase.TodoUsecase) *relay.Handler {
	s := schema.GetRootSchema()
	opts := []graphql.SchemaOpt{graphql.UseFieldResolvers()}
	parsedSchema := graphql.MustParseSchema(s, &graphqlServer{todoUsecase}, opts...)
	return &relay.Handler{Schema: parsedSchema}

}
