package model

import "time"

type Food struct {
	ID           int        `json:"id"`
	Name         string     `json:"name"`
	Description  string     `json:"description"`
	Price        int        `json:"price"`
	RestaurantID int        `json:"restaurant_id"`
	Restaurant   Restaurant `gorm:"foreignKey:RestaurantID"`
	LineFood     *LineFood  `gorm:"foreignKey:FoodID"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}
