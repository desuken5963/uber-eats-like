package repository

import "backend/internal/domain/model"

type CreateOrderRepository interface {
	FindLineFoodsByIDs(ids []int) ([]model.LineFood, error)
	CreateOrderWithLineFoods(order *model.Order, lineFoods []model.LineFood) error
}
