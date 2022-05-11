package db

import (
	"github.com/aimericsr/social-network-auth-api/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func MakeMigration(dbSource string) error {
	db, err := gorm.Open(postgres.Open(dbSource), &gorm.Config{})
	if err != nil {
		return err
	}

	db.AutoMigrate(&model.User{})
	db.Create(&model.User{Username: "admin", Password: "admin", Email: "tddest"})

	return nil
}
