package database

import (
	models "github.com/yuttasakcom/go-fiber-simple/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect()  {
	dsn := "host=localhost user=postgres password=password dbname=gosimple port=5432 sslmode=disable TimeZone=Asia/Bangkok"
	
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB = db

	db.AutoMigrate(models.User{})
}