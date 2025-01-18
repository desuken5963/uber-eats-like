package usecase

import (
	"backend/internal/domain/model"
	"backend/internal/domain/repository"
	"errors"
)

// CreateLineFoodUsecase defines the interface for creating a line food.
type CreateLineFoodUsecase interface {
	Execute(foodID int, count int) (*model.LineFood, map[string]string, error)
}

type createLineFoodUsecase struct {
	repo repository.CreateLineFoodRepository
}

// NewCreateLineFoodUsecase creates a new instance of CreateLineFoodUsecase.
func NewCreateLineFoodUsecase(repo repository.CreateLineFoodRepository) CreateLineFoodUsecase {
	return &createLineFoodUsecase{repo: repo}
}

// Execute creates a new line food or updates an existing one.
func (u *createLineFoodUsecase) Execute(foodID int, count int) (*model.LineFood, map[string]string, error) {
	food, err := u.repo.FindFoodByID(foodID)
	if err != nil {
		return nil, nil, errors.New("food not found")
	}

	// Check for active line foods in other restaurants
	activeLineFoods, err := u.repo.FindActiveByOtherRestaurant(food.RestaurantID)
	if err != nil {
		return nil, nil, err
	}

	if len(activeLineFoods) > 0 {
		return nil, map[string]string{
			"existing_restaurant": activeLineFoods[0].Restaurant.Name,
			"new_restaurant":      food.Restaurant.Name,
		}, errors.New("active line foods exist for another restaurant")
	}

	// Create or update the line food
	lineFood := &model.LineFood{
		FoodID:       food.ID,
		Count:        count,
		Active:       true,
		RestaurantID: food.RestaurantID,
	}

	err = u.repo.Save(lineFood)
	if err != nil {
		return nil, nil, err
	}

	return lineFood, nil, nil
}
