package model

import (
	"time"
)

type User struct {
	ID           uint // `gorm:"primarykey"`
	createdAt    time.Time
	updatedAt    time.Time
	email        string
	passwordHash string
}
