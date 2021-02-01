package models

func NewUser(id int) *User {
	result := &User{
		ID:          id,
		FirstName:   "Rafael",
		LastName:    "Barbosa",
		PhoneNumber: 312345678,
		Email:       "massoni@gmail.com",
	}

	return result
}
