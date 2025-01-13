package usecase

import (
	"backend/internal/domain/model"
	"backend/internal/domain/repository"
)

// RestaurantUsecase defines the business logic for restaurants.
type RestaurantUsecase interface {
	GetAllRestaurants() ([]model.Restaurant, error)
}

// restaurantUsecase implements the RestaurantUsecase interface.
type restaurantUsecase struct {
	restaurantRepo repository.RestaurantRepository
}

// NewRestaurantUsecase creates a new instance of restaurantUsecase.
func NewRestaurantUsecase(repo repository.RestaurantRepository) RestaurantUsecase {
	return &restaurantUsecase{restaurantRepo: repo}
}

// GetAllRestaurants fetches all restaurants from the repository.
func (u *restaurantUsecase) GetAllRestaurants() ([]model.Restaurant, error) {
	return u.restaurantRepo.FindAll()
}
