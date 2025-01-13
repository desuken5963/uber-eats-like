package repository

import (
	"backend/internal/domain/model"
	"backend/internal/domain/repository"

	"gorm.io/gorm"
)

type foodRepository struct {
	db *gorm.DB
}

// NewFoodRepository creates a new instance of foodRepository.
func NewFoodRepository(db *gorm.DB) repository.FoodRepository {
	return &foodRepository{db: db}
}

// FindByRestaurantID fetches all foods belonging to a specific restaurant.
func (r *foodRepository) FindByRestaurantID(restaurantID int) ([]model.Food, error) {
	var foods []model.Food
	err := r.db.Where("restaurant_id = ?", restaurantID).Find(&foods).Error
	if err != nil {
		return nil, err
	}
	return foods, nil
}
