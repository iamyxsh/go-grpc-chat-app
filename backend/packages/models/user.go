package models

import (
	"time"
)

type AuthUser struct {
	ID        uint `gorm:"primaryKey"`
	Email     string
	Password  string
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

type FriendUser struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

type ProfileUser struct {
	ID        uint `gorm:"primaryKey"`
	Email     string
	Password  string
	Name      string
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
