package repository

import (
	"backend/internal/domain/model"
	"errors"
)

var ErrNotFound = errors.New("record not found")

type ReplaceLineFoodRepository interface {
	FindActiveByOtherRestaurant(restaurantID int) ([]model.LineFood, error)
	FindFoodByID(foodID int) (*model.Food, error)
	FindLineFoodByFoodID(foodID int) (*model.LineFood, error)
	UpdateLineFoodsActiveStatus(lineFoodIDs []int, active bool) error
	Save(lineFood *model.LineFood) error
}
