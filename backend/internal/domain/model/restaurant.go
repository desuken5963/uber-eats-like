package model

import "time"

type Restaurant struct {
	ID           int
	Name         string
	Fee          int
	TimeRequired int
	Foods        []Food
	LineFoods    []LineFood
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
