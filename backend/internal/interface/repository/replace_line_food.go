package repository

import (
	"backend/internal/domain/model"
	"errors"

	"gorm.io/gorm"
)

var ErrNotFound = errors.New("record not found")

type replaceLineFoodRepository struct {
	db *gorm.DB
}

func NewReplaceLineFoodRepository(db *gorm.DB) *replaceLineFoodRepository {
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
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrNotFound
	}
	return &food, err
}

func (r *replaceLineFoodRepository) FindLineFoodByFoodID(foodID int) (*model.LineFood, error) {
	var lineFood model.LineFood
	err := r.db.Where("food_id = ?", foodID).First(&lineFood).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrNotFound
	}
	return &lineFood, err
}

func (r *replaceLineFoodRepository) UpdateLineFoodsActiveStatus(lineFoodIDs []int, active bool) error {
	if len(lineFoodIDs) == 0 {
		return nil
	}
	return r.db.Model(&model.LineFood{}).Where("id IN ?", lineFoodIDs).Update("active", active).Error
}

func (r *replaceLineFoodRepository) Save(lineFood *model.LineFood) error {
	return r.db.Save(lineFood).Error
}
