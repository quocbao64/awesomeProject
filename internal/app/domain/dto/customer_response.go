package dto

import (
	"gorm.io/gorm"
	"time"
)

type CustomerResponse struct {
	ID          uint           `json:"id"`
	Name        *string        `json:"name"`
	Address     string         `json:"address"`
	Email       string         `json:"email"`
	PhoneNumber string         `json:"phone_number"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
