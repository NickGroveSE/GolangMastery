package controllers

import (
	"03_RESTAPI/models"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserInstance struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Bio      string `json:"bio"`
	Location string `json:"location"`
}

type NullString struct {
	String string
	Valid  bool
}

func GetAllUsers(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	var users []models.User
	db.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})

}

func CreateUser(c *gin.Context) {

	var instance UserInstance
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

	db := c.MustGet("db").(*gorm.DB)
	var user models.User

	if err := db.Where("id = ?", c.Param("id")).Find(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "No User Found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})

}

func EditUser(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	var user models.User

	if err := db.Where("id = ?", c.Param("id")).Find(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "No User Found"})
		return
	}

	var instance UserInstance
	if err := c.ShouldBindJSON(&instance); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	updatingUser := models.User{
		ID:       instance.ID,
		Username: instance.Username,
		Bio:      sql.NullString{String: instance.Bio, Valid: true},
		Location: sql.NullString{String: instance.Location, Valid: true},
	}

	db.Model(&user).Updates(updatingUser)

	c.JSON(http.StatusOK, gin.H{"data": user})

}

func DeleteUser(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	var user models.User

	if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "No User Found"})
		return
	}

	db.Delete(&user)

	c.JSON(http.StatusOK, gin.H{"data": true})

}
