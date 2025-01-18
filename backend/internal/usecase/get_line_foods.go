package usecase

import (
	"backend/internal/domain/repository"
	"errors"
)

// GetLineFoodsUsecase defines the business logic for getting active line foods.
type GetLineFoodsUsecase interface {
	Execute() (map[string]interface{}, error)
}

type getLineFoodsUsecase struct {
	repo repository.GetLineFoodsRepository
}

// NewGetLineFoodsUsecase creates a new instance of GetLineFoodsUsecase.
func NewGetLineFoodsUsecase(repo repository.GetLineFoodsRepository) GetLineFoodsUsecase {
	return &getLineFoodsUsecase{repo: repo}
}

// Execute fetches all active line foods and aggregates data for the response.
func (u *getLineFoodsUsecase) Execute() (map[string]interface{}, error) {
	lineFoods, err := u.repo.FindActive()
	if err != nil {
		return nil, err
	}

	if len(lineFoods) == 0 {
		return nil, errors.New("no content")
	}

	restaurant := lineFoods[0].RestaurantID
	count := 0
	amount := 0
	lineFoodIDs := []int{}

	for _, lf := range lineFoods {
		count += lf.Count
		lineFoodIDs = append(lineFoodIDs, lf.ID)
		amount += lf.Count * 100 // Assume `price` is fixed or pre-calculated in DB
	}

	return map[string]interface{}{
		"line_food_ids": lineFoodIDs,
		"restaurant":    restaurant,
		"count":         count,
		"amount":        amount,
	}, nil
}
