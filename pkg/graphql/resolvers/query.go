package graphql

import (
	"context"
	"easyfood/pkg/graphql/gqlgen"
	"easyfood/pkg/graphql/models"
)

type queryResolver struct{}

func NewQueryResolver() gqlgen.QueryResolver {
	return new(queryResolver)
}

func (q queryResolver) Category(ctx context.Context, id int) (*models.Category, error) {
	category := models.NewCategory()
	return category, nil
}

func (q queryResolver) Dish(ctx context.Context, name string, category []string) ([]*models.Dish, error) {
	dish := models.NewDish()
	return dish, nil
}

func (q queryResolver) Restaurant(ctx context.Context, category []string) ([]*models.Restaurant, error) {
	restaurant := models.NewRestaurant()
	return restaurant, nil
}

func (q queryResolver) User(ctx context.Context, id int) (*models.User, error) {
	user := models.NewUser(id)
	return user, nil
}
