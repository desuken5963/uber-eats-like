package repository

import "backend/internal/domain/model"

type RestaurantRepository interface {
	FindAll() ([]model.Restaurant, error)
}
