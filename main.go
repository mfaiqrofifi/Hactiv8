package main

import "Task_Microservice/router"

func main() {
	PORT := ":8080"
	router.StartServer().Run(PORT)
}
