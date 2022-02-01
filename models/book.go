package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID        uint           `gorm:"primary_key;auto_increment;unique" json:"id"`
	Title     string         `gorm:"not null;default:null" json:"title"`
	Author    string         `gorm:"not nulldefault:null" json:"author"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
