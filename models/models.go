package models

import (
	"time"
)

type Task struct {
	ID          uint      `gorm:"primaryKey"`
	Description string    `gorm:"size:255;column:description"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
	Completed   bool      `gorm:"column:completed"`
	CategoryID  uint      `gorm:"column:category_id"`
	Category    Category  `gorm:"foreignKey:CategoryID"`
}

type Category struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"size:255;column:name"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
