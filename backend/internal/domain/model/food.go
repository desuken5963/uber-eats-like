package model

import "time"

type Food struct {
	ID           int
	Name         string
	Description  string
	Price        int
	RestaurantID int
	OrderID      *int
	LineFood     *LineFood
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
