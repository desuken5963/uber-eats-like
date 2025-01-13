package repository

import "backend/internal/domain/model"

// FoodRepository defines the interface for accessing food data.
type FoodRepository interface {
	FindByRestaurantID(restaurantID int) ([]model.Food, error)
}
