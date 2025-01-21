package usecase

import (
	"backend/internal/domain/model"
	"backend/internal/domain/repository"
	"errors"
)

type ReplaceLineFoodUsecase interface {
	Execute(foodID int, count int) (*model.LineFood, error)
}

type replaceLineFoodUsecase struct {
	repo repository.ReplaceLineFoodRepository
}

func NewReplaceLineFoodUsecase(repo repository.ReplaceLineFoodRepository) ReplaceLineFoodUsecase {
	return &replaceLineFoodUsecase{repo: repo}
}

func (u *replaceLineFoodUsecase) Execute(foodID int, count int) (*model.LineFood, error) {
	food, err := u.repo.FindFoodByID(foodID)
	if err != nil {
		return nil, errors.New("food not found")
	}

	activeLineFoods, err := u.repo.FindActiveByOtherRestaurant(food.RestaurantID)
	if err != nil {
		return nil, err
	}

	if len(activeLineFoods) > 0 {
		lineFoodIDs := []int{}
		for _, lf := range activeLineFoods {
			lineFoodIDs = append(lineFoodIDs, lf.ID)
		}

		err = u.repo.UpdateLineFoodsActiveStatus(lineFoodIDs, false)
		if err != nil {
			return nil, err
		}
	}

	existingLineFood, err := u.repo.FindLineFoodByFoodID(foodID)
	if err != nil && !errors.Is(err, repository.ErrNotFound) {
		return nil, err
	}

	var lineFood *model.LineFood
	if existingLineFood != nil {
		existingLineFood.Count += count
		existingLineFood.Active = true
		lineFood = existingLineFood
	} else {
		lineFood = &model.LineFood{
			FoodID:       food.ID,
			Count:        count,
			Active:       true,
			RestaurantID: food.RestaurantID,
		}
	}

	err = u.repo.Save(lineFood)
	if err != nil {
		return nil, err
	}

	return lineFood, nil
}
