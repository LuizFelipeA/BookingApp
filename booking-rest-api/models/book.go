package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	MediumPrice string         `json:"mediumPrice"`
	Author      string         `json:"author"`
	ImageURL    string         `json:"imageUrl"`
	CreatedAt   time.Time      `json:"created"`
	UpdatedAt   time.Time      `json:"updated"`
	DeletedAt   gorm.DeletedAt `json:"deleted" gorm:"index"`
}
