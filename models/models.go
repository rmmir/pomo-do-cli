package models

import "time"

type Task struct {
	ID          uint   `gorm:"primaryKey"`
	Description string `gorm:"size:255"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Completed   bool
}
