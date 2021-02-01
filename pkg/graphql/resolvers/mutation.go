package graphql

import (
	"context"
	"easyfood/pkg/graphql/gqlgen"
	"easyfood/pkg/graphql/models"
)

type mutationResolver struct{}

func NewMutationResolver() gqlgen.MutationResolver {
	return new(mutationResolver)
}

func (m mutationResolver) CreateDish(ctx context.Context, input models.CreateDishInput) (bool, error) {
	return true, nil
}

func (m mutationResolver) CreateUser(ctx context.Context, input models.CreateUserInput) (*models.User, error) {
	return models.NewUser(0), nil
}
