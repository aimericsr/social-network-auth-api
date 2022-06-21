package db

import (
	"time"

	"github.com/aimericsr/social-network-auth-api/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Open(dbSource string) (DB *gorm.DB, err error) {
	DB, err = gorm.Open(postgres.Open(dbSource), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	sqlDB, err := DB.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	return DB, nil
}

func Close(DB *gorm.DB) error {
	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func MakeMigration(dbSource string) error {
	db, err := gorm.Open(postgres.Open(dbSource), &gorm.Config{})
	if err != nil {
		return err
	}
	//defer db.close()

	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	db.Exec(`
		CREATE OR REPLACE FUNCTION function_update_timestamp()
		RETURNS TRIGGER AS $$
		BEGIN
	  		NEW.updated_at = NOW();
	  	RETURN NEW;
		END;
		$$ LANGUAGE plpgsql;
	`)

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.User{})

	//db.Create(&model.User{Username: "admin", Password: "admin", Email: "tddest"})

	//db.Create(&model.RefreshToken{Value: "tddest"})

	//user := model.User{ID: uuid.MustParse("04886b18-bd4c-4864-b4bb-55fbc7f565b2")}
	//db.Model(&user).Updates(model.User{Username: "newadmin", Password: "autrepassword"})

	//fmt.Print(user)

	return nil
}
