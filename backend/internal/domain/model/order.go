package model

import "time"

type Order struct {
	ID         int
	TotalPrice int
	LineFoods  []LineFood
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
