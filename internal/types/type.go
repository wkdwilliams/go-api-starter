package types

import (
	"time"

	"gorm.io/gorm"
)

// The abstract type for all types
type AbstractType struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"not null"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
