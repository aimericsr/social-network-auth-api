package model

import (
	"time"

	"github.com/google/uuid"
)

type Sessions struct {
	ID           uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	UserID       uuid.UUID `json:"-"`
	User         User      `gorm:"foreignKey:UserID;references:ID" json:"location,omitempty"`
	Username     string    `gorm:"type:varchar(50);not null" json:"username"`
	RefreshToken string    `gorm:"type:varchar(500);not null" json:"refresh_token"`
	UserAgent    string    `gorm:"type:varchar(50);not null" json:"user_agent"`
	ClientIP     string    `gorm:"type:varchar(50);not null" json:"client_ip"`
	IsBlocked    string    `gorm:"type:bool;not null" json:"is_blocked"`
	ExpiresAt    string    `gorm:"type:bool;not null" json:"expires_at"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}
