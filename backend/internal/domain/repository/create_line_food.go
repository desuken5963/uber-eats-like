package repository

import "backend/internal/domain/model"

// CreateLineFoodRepository defines the interface for creating a line food.
type CreateLineFoodRepository interface {
	FindActiveByOtherRestaurant(restaurantID int) ([]model.LineFood, error)
	FindFoodByID(foodID int) (*model.Food, error)
	Save(lineFood *model.LineFood) error
}
