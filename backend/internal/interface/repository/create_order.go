package repository

import (
	"backend/internal/domain/model"
	"gorm.io/gorm"
)

type createOrderRepository struct {
	db *gorm.DB
}

func NewCreateOrderRepository(db *gorm.DB) *createOrderRepository {
	return &createOrderRepository{db: db}
}

func (r *createOrderRepository) FindLineFoodsByIDs(ids []int) ([]model.LineFood, error) {
	var lineFoods []model.LineFood
	if err := r.db.Where("id IN ? AND active = ?", ids, true).Preload("Food").Preload("Restaurant").Find(&lineFoods).Error; err != nil {
		return nil, err
	}
	return lineFoods, nil
}

func (r *createOrderRepository) CreateOrderWithLineFoods(order *model.Order, lineFoods []model.LineFood) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(order).Error; err != nil {
			return err
		}

		for i := range lineFoods {
			lineFoods[i].Active = false
			lineFoods[i].OrderID = &order.ID
			if err := tx.Save(&lineFoods[i]).Error; err != nil {
				return err
			}
		}
		return nil
	})
}
