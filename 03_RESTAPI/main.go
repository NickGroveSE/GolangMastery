package main

import (
	"03_RESTAPI/models"
	"03_RESTAPI/routes"
)

func main() {

	db := models.SetupDB()
	db.AutoMigrate(&models.User{})

	r := routes.SetupRoutes(db)
	r.Run()
}
