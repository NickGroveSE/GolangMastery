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

	r.GET("/users", controllers.GetAllUsers)
	r.POST("/users", controllers.CreateUser)
	r.GET("/users/:id", controllers.GetUser)
	r.PATCH("/users/:id", controllers.EditUser)
	r.DELETE("/users/:id", controllers.DeleteUser)

	return r
}
