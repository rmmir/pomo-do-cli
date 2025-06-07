package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Task struct {
	ID          uuid.UUID     `gorm:"type:uuid;primaryKey"`
	Description string        `gorm:"size:255;column:description"`
	CreatedAt   time.Time     `gorm:"column:created_at"`
	UpdatedAt   time.Time     `gorm:"column:updated_at"`
	Completed   bool          `gorm:"column:completed"`
	WorkTime    time.Duration `gorm:"column:work_time;default:0"`
	BreakTime   time.Duration `gorm:"column:break_time;default:0"`
	CategoryID  uuid.UUID     `gorm:"type:uuid;column:category_id"`
	Category    Category      `gorm:"foreignKey:CategoryID"`
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

type Session struct {
	ID          uuid.UUID     `gorm:"type:uuid;primaryKey"`
	Name        string        `gorm:"size:255;column:name"`
	WorkTime    time.Duration `gorm:"column:work_time;default:0"`
	BreakTime   time.Duration `gorm:"column:break_time;default:0"`
	Repetitions int           `gorm:"column:repetitions;default:1"`
}

func (s *Session) BeforeCreate(tx *gorm.DB) (err error) {
	if s.ID == uuid.Nil {
		s.ID = uuid.New()
	}

	return s.Validate()
}

func (s *Session) Validate() error {
	if len(s.Name) < 3 {
		return errors.New("session name must be at least 3 characters long")
	}

	if s.WorkTime < 15*time.Minute {
		return errors.New("work time must be at least 15 minutes")
	}

	if s.BreakTime < 5*time.Minute {
		return errors.New("break time must be at least 5 minutes")
	}

	if s.Repetitions < 1 {
		return errors.New("repetitions must be at least 1")
	}

	return nil
}
