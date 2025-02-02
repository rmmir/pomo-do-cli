package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Task struct {
	ID          uint   `gorm:"primaryKey"`
	Description string `gorm:"size:255"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Completed   bool
	CategoryID  uint
	Category    Category `gorm:"foreignKey:CategoryID"`
}

type Category struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (t *Task) BeforeSave(tx *gorm.DB) (err error) {
	if len(t.Description) < 3 {
		return errors.New("task description must be at least 3 characters long")
	}

	return nil
}
