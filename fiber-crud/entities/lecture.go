package entities

import (
	"gorm.io/gorm"
	"time"
)

type Lecture struct {
	Id        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"  json:"deleted_at"`
	Name      string         `json:"name"`
}
