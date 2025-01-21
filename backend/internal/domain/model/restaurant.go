package model

import "time"

type Restaurant struct {
	ID           int        `json:"id"`
	Name         string     `json:"name"`
	Fee          int        `json:"fee"`
	TimeRequired int        `json:"time_required"`
	Foods        []Food     `json:"foods"`
	LineFoods    []LineFood `json:"line_foods"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}
