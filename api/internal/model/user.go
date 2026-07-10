package model

import (
	"time"
)

type User struct {
	Id          uint      `gorm:"primaryKey" json:"id"`
	FirebaseUID string    `gorm:"uniqueIndex;size:128;default:null" json:"-"`
	Name        string    `gorm:"size:100;not null" json:"name"`
	Email       string    `gorm:"size:100;unique;not null" json:"email"`
	Password    string    `gorm:"not null" json:"-"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
