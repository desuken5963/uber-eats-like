package model

import "time"

type LineFood struct {
	ID           int
	FoodID       int
	RestaurantID int
	OrderID      *int
	Count        int
	Active       bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
