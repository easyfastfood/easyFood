package graphql

import "easyfood/pkg/graphql/gqlgen"

type app struct{}

func NewResolverRoot() gqlgen.ResolverRoot {
	return new(app)
}

func (a app) Query() gqlgen.QueryResolver {
	return NewQueryResolver()
}
