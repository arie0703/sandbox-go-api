package main

import (
	controller "sandbox-go-api/controllers"
	"sandbox-go-api/database"
)

func main() {

	database.Init()
	database.GetDB()
	r := controller.GetRouter()
	r.Run()
}
