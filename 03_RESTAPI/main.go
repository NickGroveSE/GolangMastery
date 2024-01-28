package main

import (
	"03_RESTAPI/models"
	"03_RESTAPI/routes"
)

func main() {

	db := models.SetupDB()
	db.AutoMigrate(&models.User{}, &models.Brewery{})

	r := routes.SetupRoutes(db)
	r.Run()
}
