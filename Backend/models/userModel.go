package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID       string `gorm:"type:char(36);default:uuid()"` 
	Nama string `json:"Nama" binding:"required"`
	Password string `json:"Password" binding:"required"`
	Email string `json:"Email" binding:"required"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String()
	return
}