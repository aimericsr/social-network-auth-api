package model

import (
	"github.com/google/uuid"
)

type Artist struct {
	ID        uuid.UUID   `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	Album     string      `gorm:"type:varchar(200);not null" json:"libelle"`
	FirstName string      `gorm:"type:varchar(20);not null" json:"first_name"`
	LastName  string      `gorm:"type:varchar(20);not null" json:"last_name"`
	Festival  []*Festival `gorm:"many2many:artists_festivals;" json:"festivals"`
}
