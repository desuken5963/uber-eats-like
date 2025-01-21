package model

import "time"

type LineFood struct {
	ID           int        `json:"id"`
	FoodID       int        `json:"food_id"`
	Food         Food       `gorm:"foreignKey:FoodID"`
	RestaurantID int        `json:"restaurant_id"`
	Restaurant   Restaurant `gorm:"foreignKey:RestaurantID"`
	OrderID      *int       `json:"order_id"`
	Order        *Order     `gorm:"foreignKey:OrderID"`
	Count        int        `json:"count"`
	Active       bool       `json:"active"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}
