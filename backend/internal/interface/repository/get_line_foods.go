package repository

import (
	"backend/internal/domain/model"
	"backend/internal/domain/repository"

	"gorm.io/gorm"
)

type getLineFoodsRepository struct {
	db *gorm.DB
}

// NewGetLineFoodsRepository creates a new instance of getLineFoodsRepository.
func NewGetLineFoodsRepository(db *gorm.DB) repository.GetLineFoodsRepository {
	return &getLineFoodsRepository{db: db}
}

// FindActive fetches all active line foods from the database.
func (r *getLineFoodsRepository) FindActive() ([]model.LineFood, error) {
	var lineFoods []model.LineFood
	err := r.db.Where("active = ?", true).Preload("Restaurant").Find(&lineFoods).Error
	if err != nil {
		return nil, err
	}
	return lineFoods, nil
}
