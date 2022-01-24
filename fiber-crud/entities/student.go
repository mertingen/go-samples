package entities

import (
	"gorm.io/gorm"
	"time"
)

type Student struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"  json:"deleted_at"`
	Fullname  string         `json:"fullname"`
	Email     string         `json:"email"`
	Age       int            `json:"age"`
}
