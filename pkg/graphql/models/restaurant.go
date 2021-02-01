package models

import "time"

func NewRestaurant() []*Restaurant {
	result := make([]*Restaurant, 0)
	result = append(result, &Restaurant{
		ID:          0,
		Category:    &Category{},
		OpenHour:    time.Now(),
		CloseHour:   time.Now(),
		OpenDays:    AllWeekdays,
		City:        &City{},
		Name:        "Restaurante do Zé",
		Description: nil,
		PhoneNumber: "3197621337",
		Address:     "Rua do Zé, 991",
	})

	return result
}
