package repository

import (
	"backend/internal/domain/model"
	"backend/internal/domain/repository"

	"gorm.io/gorm"
)

// restaurantRepository implements the RestaurantRepository interface.
type restaurantRepository struct {
	db *gorm.DB
}

// NewRestaurantRepository creates a new instance of restaurantRepository.
func NewRestaurantRepository(db *gorm.DB) repository.RestaurantRepository {
	return &restaurantRepository{db: db}
}

// FindAll fetches all restaurants from the database.
func (r *restaurantRepository) FindAll() ([]model.Restaurant, error) {
	var restaurants []model.Restaurant
	err := r.db.Preload("Foods").Preload("LineFoods").Find(&restaurants).Error
	if err != nil {
		return nil, err
	}
	return restaurants, nil
}
