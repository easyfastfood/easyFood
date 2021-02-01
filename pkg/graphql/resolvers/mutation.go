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

func (m mutationResolver) CreateCategory(ctx context.Context, name string) (bool, error) {
	return true, nil
}

func (m mutationResolver) CreateRestaurant(ctx context.Context, input models.CreateRestaurantInput) (*models.Restaurant, error) {
	return models.NewRestaurant()[0], nil
}
