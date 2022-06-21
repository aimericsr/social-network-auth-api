package model

import (
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	Username  string    `gorm:"type:varchar(20);unique;not null" json:"username"`
	FirstName string    `gorm:"type:varchar(20);not null" json:"first_name"`
	LastName  string    `gorm:"type:varchar(20);not null" json:"last_name"`
	Password  string    `gorm:"type:varchar(100);not null" json:"password"`
	Email     string    `gorm:"type:varchar(50);unique;not null" json:"email"`
}
