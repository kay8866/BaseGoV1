package models

import (
	"time"
)

// User represents the user model
type User struct {
	ID        uint       `gorm:"autoID" json:"-"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"-"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"-"`
	DeletedAt *time.Time `gorm:"autoDelete" json:"-"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
}
