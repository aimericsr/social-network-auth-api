package model

import (
	"time"

	"github.com/google/uuid"
)

type Festival struct {
	ID      uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	Date    time.Time `json:"created_at"`
	Libelle string    `gorm:"type:varchar(200);not null" json:"libelle"`
	Artist  []*Artist `gorm:"many2many:artists_festivals;" json:"artists"`
}