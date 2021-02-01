package models

func NewUser(id int) *User {
	result := &User{
		ID:          id,
		FirstName:   "Rafael",
		LastName:    "Barbosa",
		PhoneNumber: "3199827382",
		Email:       "massoni@gmail.com",
	}

	return result
}
