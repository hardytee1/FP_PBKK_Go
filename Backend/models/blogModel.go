package models

import (
	"gorm.io/gorm"
)

type Blog struct {
	gorm.Model 
	Content string `json:"content" binding:"required"`
	Caption string `json:"caption" binding:"required"`
	UserID string `json:"user_id" binding:"required"`
}