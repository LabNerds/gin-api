package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"gin-api/config"

	"gin-api/app/models"
)

var db *gorm.DB
var err error

func Connect() {
	c := config.GetConfig()
	db, err = gorm.Open(postgres.New(
		postgres.Config{
			DSN: fmt.Sprintf(
				"user=%s password=%s dbname=%s port=%s sslmode=%s",
				c.GetString("db.user"),
				c.GetString("db.password"),
				c.GetString("db.dbname"),
				c.GetString("db.port"),
				c.GetString("db.sllmode"),
			),
			PreferSimpleProtocol: true,
		}), &gorm.Config{})

	if err == nil {
		fmt.Println("DB connected")
		setupDatabaseConection()
	}
}

func setupDatabaseConection() {
	db.AutoMigrate(&models.User{})
}

func GetDB() *gorm.DB {
	if err != nil {
		panic("Failed to get db driver")
	}

	return db
}

func CloseDBConnection() {
	current_db, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}
	current_db.Close()
}
