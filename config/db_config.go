package config

import (
	"os"
	"pawb_backend/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func ConnectDB() {
	var dsn string
	dbEnv := os.Getenv("DB_ENVIRONMENT")
	if dbEnv == "dev" {
		dbUser := os.Getenv("DB_USERNAME")
		dbPass := os.Getenv("DB_PASSWORD")
		dbHost := os.Getenv("DB_HOST")
		dbPort := os.Getenv("DB_PORT")
		dbDatabase := os.Getenv("DB_DATABASE")

		dsn = dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbDatabase + "?charset=utf8mb4&parseTime=True&loc=Local"
	} else if dbEnv == "prod" {
		dsn = os.Getenv("DSN")
	}
	var err error
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic("Failed to connect to database!")
	}

	err = Db.AutoMigrate(&entity.User{}, &entity.Role{}, &entity.BodyPart{}, &entity.PsoriasisImage{}, &entity.PatientClinician{}, &entity.Feedback{})
	if err != nil {
		panic("Failed to migrate database!")
	}
}

func CloseDb() {
	db, err := Db.DB()
	if err != nil {
		panic("Failed to close database!")
	}
	err = db.Close()
	if err != nil {
		return
	}
}
