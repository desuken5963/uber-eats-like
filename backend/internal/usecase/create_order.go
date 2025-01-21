package usecase

import (
	"backend/internal/domain/model"
	"backend/internal/domain/repository"
	"errors"
)

type CreateOrderUsecase interface {
	Execute(lineFoodIDs []int) error
}

type createOrderUsecase struct {
	repo repository.CreateOrderRepository
}

func NewCreateOrderUsecase(repo repository.CreateOrderRepository) CreateOrderUsecase {
	return &createOrderUsecase{repo: repo}
}

func (u *createOrderUsecase) Execute(lineFoodIDs []int) error {
	lineFoods, err := u.repo.FindLineFoodsByIDs(lineFoodIDs)
	if err != nil || len(lineFoods) == 0 {
		return errors.New("invalid line food IDs")
	}

	totalPrice := calculateTotalPrice(lineFoods)
	order := &model.Order{
		TotalPrice: totalPrice,
		LineFoods:  lineFoods,
	}

	if err := u.repo.CreateOrderWithLineFoods(order, lineFoods); err != nil {
		return err
	}

	return nil
}

func calculateTotalPrice(lineFoods []model.LineFood) int {
	total := 0
	for _, lineFood := range lineFoods {
		total += lineFood.Count * lineFood.Food.Price
	}
	if len(lineFoods) > 0 {
		total += lineFoods[0].Restaurant.Fee
	}
	return total
}
