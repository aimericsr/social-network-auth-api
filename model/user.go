package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Username  string    `gorm:"type:varchar(20);unique;not null"`
	Password  string    `gorm:"type:varchar(20);not null"`
	Email     string    `gorm:"type:varchar(50);unique;not null"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
