package usecase

import (
	"backend/internal/domain/model"
	"backend/internal/domain/repository"
	"errors"
)

// ReplaceLineFoodUsecase defines the interface for replacing line foods.
type ReplaceLineFoodUsecase interface {
	Execute(foodID int, count int) (*model.LineFood, error)
}

type replaceLineFoodUsecase struct {
	repo repository.ReplaceLineFoodRepository
}

// NewReplaceLineFoodUsecase creates a new instance of ReplaceLineFoodUsecase.
func NewReplaceLineFoodUsecase(repo repository.ReplaceLineFoodRepository) ReplaceLineFoodUsecase {
	return &replaceLineFoodUsecase{repo: repo}
}

// Execute replaces active line foods with a new one for the given food.
func (u *replaceLineFoodUsecase) Execute(foodID int, count int) (*model.LineFood, error) {
	// Find the food by ID
	food, err := u.repo.FindFoodByID(foodID)
	if err != nil {
		return nil, errors.New("food not found")
	}

	// Deactivate active line foods from other restaurants
	activeLineFoods, err := u.repo.FindActiveByOtherRestaurant(food.RestaurantID)
	if err != nil {
		return nil, err
	}

	lineFoodIDs := []int{}
	for _, lf := range activeLineFoods {
		lineFoodIDs = append(lineFoodIDs, lf.ID)
	}

	err = u.repo.UpdateLineFoodsActiveStatus(lineFoodIDs, false)
	if err != nil {
		return nil, err
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
		return nil, err
	}

	return lineFood, nil
}
