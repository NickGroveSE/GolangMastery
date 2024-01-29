package controllers

import (
	"03_RESTAPI/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BreweryInstance struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	BreweryType string `json:"type"`
	Location    string `json:"location"`
}

func GetAllBreweries(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	var breweries []models.Brewery
	db.Find(&breweries)

	c.JSON(http.StatusOK, gin.H{"data": breweries})

}

func CreateBrewery(c *gin.Context) {

	var instance BreweryInstance
	if err := c.ShouldBindJSON(&instance); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	brewery := models.Brewery{
		ID:          instance.ID,
		Name:        instance.Name,
		BreweryType: instance.BreweryType,
		Location:    instance.Location,
	}

	db := c.MustGet("db").(*gorm.DB)
	db.Create(&brewery)

	c.JSON(http.StatusOK, gin.H{"data": brewery})

}

func GetBrewery(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	var brewery models.Brewery

	if err := db.Where("id = ?", c.Param("id")).Find(&brewery).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "No Brewery Found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": brewery})

}

func EditBrewery(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	var brewery models.Brewery

	if err := db.Where("id = ?", c.Param("id")).Find(&brewery).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "No Brewery Found"})
		return
	}

	var instance BreweryInstance
	if err := c.ShouldBindJSON(&instance); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	updatingBrewery := models.Brewery{
		ID:          instance.ID,
		Name:        instance.Name,
		BreweryType: instance.BreweryType,
		Location:    instance.Location,
	}

	db.Model(&brewery).Updates(updatingBrewery)

	c.JSON(http.StatusOK, gin.H{"data": brewery})

}

func DeleteBrewery(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	var brewery models.Brewery

	if err := db.Where("id = ?", c.Param("id")).First(&brewery).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "No Brewery Found"})
		return
	}

	db.Delete(&brewery)

	c.JSON(http.StatusOK, gin.H{"data": true})

}
