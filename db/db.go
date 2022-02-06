package db

import (
	"fmt"

	"gin-api/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func Init() {
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
	}
}

func GetDB() *gorm.DB {
	return db
}
