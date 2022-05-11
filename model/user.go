package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type User struct {
	ID        uuid.UUID `gorm:"primary_key"`
	Username  string    `gorm:"type:varchar(20);unique;not null"`
	Password  string    `gorm:"type:varchar(20);not null"`
	Email     string    `gorm:"type:varchar(50); unique;not null"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}

func (u *User) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("ID", uuid.New())
}
