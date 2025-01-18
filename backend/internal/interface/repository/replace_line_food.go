package repository

import (
	"backend/internal/domain/model"
	"backend/internal/domain/repository"

	"gorm.io/gorm"
)

type replaceLineFoodRepository struct {
	db *gorm.DB
}

// NewReplaceLineFoodRepository creates a new instance of replaceLineFoodRepository.
func NewReplaceLineFoodRepository(db *gorm.DB) repository.ReplaceLineFoodRepository {
	return &replaceLineFoodRepository{db: db}
}

func (r *replaceLineFoodRepository) FindActiveByOtherRestaurant(restaurantID int) ([]model.LineFood, error) {
	var lineFoods []model.LineFood
	err := r.db.Where("active = ? AND restaurant_id != ?", true, restaurantID).Find(&lineFoods).Error
	if err != nil {
		return nil, err
	}
	return lineFoods, nil
}

func (r *replaceLineFoodRepository) FindFoodByID(foodID int) (*model.Food, error) {
	var food model.Food
	err := r.db.Preload("Restaurant").First(&food, foodID).Error
	if err != nil {
		return nil, err
	}
	return &food, nil
}

func (r *replaceLineFoodRepository) UpdateLineFoodsActiveStatus(lineFoodIDs []int, active bool) error {
	return r.db.Model(&model.LineFood{}).Where("id IN ?", lineFoodIDs).Update("active", active).Error
}

func (r *replaceLineFoodRepository) Save(lineFood *model.LineFood) error {
	return r.db.Save(lineFood).Error
}
