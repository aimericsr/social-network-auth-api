package db

import (
	"fmt"
	"time"

	"github.com/aimericsr/social-network-auth-api/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

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

func Close() error {
	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func MakeMigration(dbSource string) error {
	var db *gorm.DB
	var err error
	for i := 1; i <= 5; i++ {
		time.Sleep(time.Millisecond * 2000)
		fmt.Print("Try to make migration \n")
		db, err = gorm.Open(postgres.Open(dbSource), &gorm.Config{})
		if err != nil {
			//return err
		} else {
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
			db.AutoMigrate(&model.Festival{}, &model.Artist{})

			// user := model.User{Username: "aimeric-sr", FirstName: "aimeric", LastName: "sorin", Password: "jesuissecret", Email: "aimeric.sorin@gmail.com"}
			// db.Create(&user)

			// artist := model.Artist{Album: "UMLA", FirstName: "Alpha", LastName: "Wann"}
			// db.Create(&artist)

			// fesival := model.Festival{Date: time.Now(), Libelle: "helfest"}
			// db.Create(&fesival)

			//db.Create(&model.User{Username: "admin", Password: "admin", Email: "tddest"})

			//db.Create(&model.RefreshToken{Value: "tddest"})

			//user := model.User{ID: uuid.MustParse("04886b18-bd4c-4864-b4bb-55fbc7f565b2")}
			//db.Model(&user).Updates(model.User{Username: "newadmin", Password: "autrepassword"})

			//fmt.Print(user)

			return nil
		}
	}
	return err

}
