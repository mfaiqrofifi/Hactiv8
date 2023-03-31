package main

import (
	"ShowCase/config"

	"ShowCase/app"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
	err := config.InitPostgrees()
	if err != nil {
		panic(err)
	}
}

func main() {
	app.StartApplication()
}
