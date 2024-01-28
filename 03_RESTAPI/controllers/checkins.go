package controllers

import (
	"03_RESTAPI/models"

	"github.com/gin-gonic/gin"
)

type CheckInInstance struct {
	ID          uint    `json:"id"`
	UserID      uint    `json:"user"`
	BeerID      uint    `json:"beer"`
	Rating      float64 `json:"rating"`
	Description string  `json:"description"`
	User        models.User
	Beer        models.Beer
}

func AllCheckIns(c *gin.Context) {

	// Code

}

func PostCheckIn(c *gin.Context) {

	// Brewery & Beer Routes Needed First!

	// db := c.MustGet("db").(*gorm.DB)
	// var user models.User

	// if err := db.Where("id = ?", c.Param("id")).Find(&user).Error; err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"Error": "No User Found"})
	// 	return
	// }

	// var instance CheckInInstance
	// if err := c.ShouldBindJSON(&instance); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// checkIn := models.CheckIn{
	// 	ID:     instance.ID,
	// 	UserID: instance.UserID,
	// 	BeerID: instance.BeerID,
	// 	Rating: sql.NullFloat64{Float64: instance.Rating, Valid: true},
	// }

}

func QueryCheckIns(c *gin.Context) {

	// Code

}

func EditCheckIn(c *gin.Context) {

	// Code

}

func DeleteCheckIn(c *gin.Context) {

	// Code

}
