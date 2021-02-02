package services

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"

	"easyfood/pkg/entity"
)

type RestaurantService interface {
	Get(ctx context.Context, id *int) ([]*entity.Restaurant, error)
	GetByCategory(ctx context.Context, categoryID int) ([]*entity.Restaurant, error)
	GetByDish(ctx context.Context, dishID int) (*entity.Restaurant, error)
}

type restaurantService struct {
	db *sqlx.DB
}

func NewRestaurantService(db *sqlx.DB) RestaurantService {
	return restaurantService{db: db}
}

func (d restaurantService) Get(ctx context.Context, id *int) ([]*entity.Restaurant, error) {
	result := make([]*entity.Restaurant, 0)

	query := `
		SELECT * FROM restaurantes
	`

	if id != nil {
		query = fmt.Sprintf(`
			SELECT * FROM restaurantes WHERE id = %d
		`, *id)
	}

	err := d.db.Select(&result, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return result, nil
}

func (d restaurantService) GetByCategory(ctx context.Context, categoryID int) ([]*entity.Restaurant, error) {
	result := make([]*entity.Restaurant, 0)

	query := `
		SELECT r.id, r.id_horario_abertura, r.id_horario_fechamento, r.id_cidade, 
		r.dias_funcionamento, r.nome, r.descricao, r.telefone, r.endereco 
		FROM restaurantes r 
		INNER JOIN restaurante-categoria rc ON rc.id_restaurante = r.id
		WHERE rc.id_categoria = ?
	`

	err := d.db.Select(&result, query, categoryID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return result, nil
}

func (d restaurantService) GetByDish(ctx context.Context, dishID int) (*entity.Restaurant, error) {
	result := new(entity.Restaurant)

	query := `
		SELECT r.id, r.id_horario_abertura, r.id_horario_fechamento, r.id_cidade, 
		r.dias_funcionamento, r.nome, r.descricao, r.telefone, r.endereco 
		FROM restaurantes r INNER JOIN pratos p ON p.id_restaurante = r.id
		WHERE p.id = ?
	`

	err := d.db.Get(result, query, dishID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return result, nil
}
