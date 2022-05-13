package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	Username    string    `gorm:"type:varchar(20);unique;not null" json:"username"`
	Password    string    `gorm:"type:varchar(100);not null" json:"password"`
	Email       string    `gorm:"type:varchar(50);unique;not null" json:"email"`
	Description string    `gorm:"type:varchar(500);unique;not null" json:"description"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
