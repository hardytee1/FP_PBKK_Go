package models

import (
	"gorm.io/gorm"
	"time"
)

type Blog struct {
	gorm.Model 
	Content string `json:"content" binding:"required"`
	Caption string `json:"caption" binding:"required"`
	UserID string `json:"user_id" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}