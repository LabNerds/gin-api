package main

import (
	"flag"
	"fmt"
	"os"

	"gin-api/config"
	"gin-api/db"
	"gin-api/server"
)

type User struct {
	Name string
}

func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()

	config.Init(*environment)
	db.Init()
	database := db.GetDB()

	database.Migrator().CreateTable(&User{})
	server.Init()
}
