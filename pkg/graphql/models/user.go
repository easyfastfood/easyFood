package models

import "easyfood/pkg/entity"

func NewUser(u entity.User) *User {
	if u.IsEmpty() {
		return nil
	}

	result := &User{
		ID:          u.Id,
		FirstName:   u.FirstName,
		LastName:    u.LastName,
		PhoneNumber: u.PhoneNumber,
		Email:       u.Email,
	}

	return result
}
