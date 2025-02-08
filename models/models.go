package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Task struct {
    ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
    Description string    `gorm:"size:255;column:description"`
    CreatedAt   time.Time `gorm:"column:created_at"`
    UpdatedAt   time.Time `gorm:"column:updated_at"`
    Completed   bool      `gorm:"column:completed"`
    CategoryID  uuid.UUID `gorm:"type:uuid;column:category_id"`
    Category    Category  `gorm:"foreignKey:CategoryID"`
}

func (t *Task) BeforeCreate(tx *gorm.DB) (err error) {
    if t.ID == uuid.Nil {
        t.ID = uuid.New()
    }

    return t.Validate()
}

func (t *Task) Validate() error {
    if len(t.Description) < 3 {
        return errors.New("task description must be at least 3 characters long")
    }

    return nil
 }

type Category struct {
    ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
    Name      string    `gorm:"size:255;column:name"`
    CreatedAt time.Time `gorm:"column:created_at"`
    UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (c *Category) BeforeCreate(tx *gorm.DB) (err error) {
    if c.ID == uuid.Nil {
        c.ID = uuid.New()
    }

    return c.Validate()
}

func (c *Category) Validate() error {
    if len(c.Name) < 3 {
        return errors.New("category name must be at least 3 characters long")
    }

    return nil
 }
