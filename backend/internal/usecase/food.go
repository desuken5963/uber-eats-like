package usecase

import (
	"backend/internal/domain/model"
	"backend/internal/domain/repository"
)

// FoodUsecase defines the business logic for foods.
type FoodUsecase interface {
	GetFoodsByRestaurantID(restaurantID int) ([]model.Food, error)
}

type foodUsecase struct {
	foodRepo repository.FoodRepository
}

// NewFoodUsecase creates a new instance of foodUsecase.
func NewFoodUsecase(repo repository.FoodRepository) FoodUsecase {
	return &foodUsecase{foodRepo: repo}
}

// GetFoodsByRestaurantID fetches foods by restaurant ID.
func (u *foodUsecase) GetFoodsByRestaurantID(restaurantID int) ([]model.Food, error) {
	return u.foodRepo.FindByRestaurantID(restaurantID)
}
