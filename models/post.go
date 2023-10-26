package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	ID      uint `gorm:"primaryKey"`
	Title   string
	Content string
}
