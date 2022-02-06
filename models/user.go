package models

import "gin-api/db"

type User struct {
	Name string
	Email string
}

func Run_User_Migrate() {
	database := db.GetDB()
	database.Migrator().AutoMigrate(&User{})
}
