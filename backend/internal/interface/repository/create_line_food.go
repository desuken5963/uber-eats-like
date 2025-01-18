package repository

import (
	"backend/internal/domain/model"
	"backend/internal/domain/repository"

	"gorm.io/gorm"
)

type createLineFoodRepository struct {
	db *gorm.DB
}

// NewCreateLineFoodRepository creates a new instance of createLineFoodRepository.
func NewCreateLineFoodRepository(db *gorm.DB) repository.CreateLineFoodRepository {
	return &createLineFoodRepository{db: db}
}

func (r *createLineFoodRepository) FindActiveByOtherRestaurant(restaurantID int) ([]model.LineFood, error) {
	var lineFoods []model.LineFood
	err := r.db.Where("active = ? AND restaurant_id != ?", true, restaurantID).Find(&lineFoods).Error
	if err != nil {
		return nil, err
	}
	return lineFoods, nil
}

func (r *createLineFoodRepository) FindFoodByID(foodID int) (*model.Food, error) {
	var food model.Food
	err := r.db.First(&food, foodID).Error
	if err != nil {
		return nil, err
	}
	return &food, nil
}

func (r *createLineFoodRepository) Save(lineFood *model.LineFood) error {
	return r.db.Save(lineFood).Error
}
