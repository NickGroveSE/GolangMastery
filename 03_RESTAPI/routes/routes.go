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
	r.POST("/users", controllers.CreateUser)
	return r
}
