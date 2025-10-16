package models

import (
	"time"

	"gorm.io/gorm"
)

// Todo represents a to-do item
type Todo struct {
    ID          uint           `gorm:"primaryKey" json:"id"`
    Title       string         `gorm:"type:text;not null" json:"title"`
    Description string         `gorm:"type:text" json:"description,omitempty"`
    Completed   bool           `gorm:"default:false;not null" json:"completed"`
    CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
    DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
