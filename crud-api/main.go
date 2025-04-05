package main

import (
	controller "sandbox-go-api/crud-api/controllers"
	"sandbox-go-api/crud-api/database"
)

func main() {

	database.Init()
	database.GetDB()
	r := controller.GetRouter()
	r.Run()
}
