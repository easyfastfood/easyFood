package models

import (
	"easyfood/pkg/entity"
)

func NewRestaurant(restaurants ...*entity.Restaurant) []*Restaurant {
	result := make([]*Restaurant, 0)
	for _, restaurant := range restaurants {
		result = append(result, &Restaurant{
			ID:          restaurant.Id,
			Name:        restaurant.Name,
			Description: restaurant.Description,
			PhoneNumber: restaurant.PhoneNumber,
			Address:     restaurant.Address,
		})
	}

	return result
}
