package models

import (
	"time"
)

type FriendshipStatus string

const (
	SENT     FriendshipStatus = "sent"
	ACCEPTED FriendshipStatus = "accepted"
	REJECTED FriendshipStatus = "rejected"
)

type Friendship struct {
	ID        uint `gorm:"primaryKey"`
	From      FriendUser
	FromID    uint
	To        FriendUser
	ToID      uint
	Status    FriendshipStatus `gorm:"type:friendship_status"`
	CreatedAt time.Time        `gorm:"autoCreateTime"`
}
