package repository

import "backend/internal/domain/model"

// GetLineFoodsRepository defines the interface for fetching active line foods.
type GetLineFoodsRepository interface {
	FindActive() ([]model.LineFood, error)
}
