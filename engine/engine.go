package engine

import (
	"context"

	"github.com/99designs/gqlgen/graphql"

	"github.com/vektah/gqlparser"
	"github.com/vektah/gqlparser/ast"
)

func NewEngine(schemaDefinition string) graphql.ExecutableSchema {
	schema := gqlparser.MustLoadSchema(&ast.Source{Input: schemaDefinition, Name: "schema.graphql"})

	return &engine{
		schema: schema,
	}
}

type engine struct {
	schema *ast.Schema
}

func (e *engine) Schema() *ast.Schema {
	return e.schema
}

func (e *engine) Complexity(typeName, fieldName string, childComplexity int, args map[string]interface{}) (int, bool) {
	// TODO
	return 1, true
}

func (e *engine) Query(ctx context.Context, op *ast.OperationDefinition) *graphql.Response {
	return &graphql.Response{
		Data: []byte("{}"),
	}
}

func (e *engine) Mutation(ctx context.Context, op *ast.OperationDefinition) *graphql.Response {
	return &graphql.Response{
		Data: []byte("{}"),
	}

}

func (e *engine) Subscription(ctx context.Context, op *ast.OperationDefinition) func() *graphql.Response {
	return nil
}
