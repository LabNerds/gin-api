package main

import (
	"flag"
	"fmt"
	"os"

	"gin-api/config"
	"gin-api/db"
	"gin-api/server"
)

func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()

	config.Init(*environment)

	defer db.CloseDBConnection()
	db.Connect()

	server.Init()
}
