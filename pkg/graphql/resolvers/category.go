package graphql

import (
	"context"
	"easyfood/pkg/graphql/gqlgen"
	"easyfood/pkg/graphql/models"
	"easyfood/services"
)

type categoryResolver struct {
	services services.All
}

func NewCategoryResolver(s services.All) gqlgen.CategoryResolver {
	return categoryResolver{services: s}
}

func (c categoryResolver) Dishes(ctx context.Context, category *models.Category) ([]*models.Dish, error) {
	dishes, err := c.services.Dish.GetByCategory(ctx, category.ID)
	if err != nil {
		return nil, err
	}
	if len(dishes) == 0 {
		return nil, nil
	}
	return models.NewDish(dishes...), nil
}
