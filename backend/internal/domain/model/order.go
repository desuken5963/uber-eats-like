package model

import "time"

type Order struct {
	ID         int        `json:"id"`
	TotalPrice int        `json:"total_price"`
	LineFoods  []LineFood `json:"line_foods"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}
