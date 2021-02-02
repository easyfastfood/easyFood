package services

import (
	"context"

	"github.com/jmoiron/sqlx"

	"easyfood/pkg/entity"
)

type UserService interface {
	Get(ctx context.Context, id int) (*entity.User, error)
}

type userService struct {
	db *sqlx.DB
}

func NewUserService(db *sqlx.DB) UserService {
	return userService{db: db}
}

func (u userService) Get(ctx context.Context, id int) (*entity.User, error) {
	result := new(entity.User)
	query := "SELECT * FROM usuarios WHERE id = ?"

	if err := u.db.Get(&result, query, id); err != nil {
		return nil, err
	}

	return result, nil
}