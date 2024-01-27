package controllers

import (
	"03_RESTAPI/models"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CreateUserInstance struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Bio      string `json:"bio"`
	Location string `json:"location"`
}

type NullString struct {
	String string
	Valid  bool
}

func CreateUser(c *gin.Context) {

	var instance CreateUserInstance
	if err := c.ShouldBindJSON(&instance); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		ID:       instance.ID,
		Username: instance.Username,
		Bio:      sql.NullString{String: instance.Bio, Valid: true},
		Location: sql.NullString{String: instance.Location, Valid: true},
	}

	db := c.MustGet("db").(*gorm.DB)
	db.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})

}

func GetUser(c *gin.Context) {

	// Code

}

func EditUser(c *gin.Context) {

	// Code

}

func DeleteUser(c *gin.Context) {

	// Code

}
