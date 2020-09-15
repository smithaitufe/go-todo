package graphql

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/smithaitufe/go-todo/transport/graphql/resolver"
	"github.com/smithaitufe/go-todo/transport/graphql/schema"
	"github.com/smithaitufe/go-todo/usecase"
)

func NewGraphQLServer(todoUsecase usecase.TodoUsecase, userUsecase usecase.UserUsecase) *relay.Handler {
	s := schema.GetRootSchema()

	r := resolver.NewRootResolver(todoUsecase, userUsecase)

	opts := []graphql.SchemaOpt{graphql.UseFieldResolvers()}
	parsedSchema := graphql.MustParseSchema(s, &r, opts...)
	return &relay.Handler{Schema: parsedSchema}

}
