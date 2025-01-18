package repository

import "backend/internal/domain/model"

// ReplaceLineFoodRepository defines the interface for replacing line foods.
type ReplaceLineFoodRepository interface {
	FindActiveByOtherRestaurant(restaurantID int) ([]model.LineFood, error)
	FindFoodByID(foodID int) (*model.Food, error)
	UpdateLineFoodsActiveStatus(lineFoodIDs []int, active bool) error
	Save(lineFood *model.LineFood) error
}
