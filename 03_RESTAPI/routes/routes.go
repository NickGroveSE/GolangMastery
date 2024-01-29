package routes

import (
	"03_RESTAPI/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	// Test Routes
	r.GET("/user/all", controllers.GetAllUsers)
	r.GET("/brewery/all", controllers.GetAllBreweries)

	// User Routes
	r.POST("/user", controllers.CreateUser)
	r.GET("/user/:id", controllers.GetUser)
	r.PATCH("/user/:id", controllers.EditUser)
	r.DELETE("/user/:id", controllers.DeleteUser)

	// Brewery Routes
	r.POST("/brewery", controllers.CreateBrewery)
	r.GET("/brewery/:id", controllers.GetBrewery)
	r.PATCH("/brewery/:id", controllers.EditBrewery)

	return r
}
