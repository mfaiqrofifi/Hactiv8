package main

import (
	"DDD/config"

	"DDD/app"

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
