package controllers

import (
	"net/http"
	"strconv"
	// "fmt"
	"github.com/gin-gonic/gin"
	"go-heartbeat-job/database"
	"go-heartbeat-job/models"
)

func WeatherGet(c *gin.Context) {
	db 			:= database.GetDB()
	id, _ 		:= strconv.Atoi(c.Param("id"))
	Weathers 	:= models.Weathers{}
	err 		:= db.Find(&Weathers, uint(id)).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Weathers)
}

func WeatherGetList(c *gin.Context) {
	db := database.GetDB()
	Weathers := []models.Weathers{}
	err := db.Find(&Weathers).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Weathers)
}

func WeatherCreate(c *gin.Context) {
	db := database.GetDB()

	var input models.Weathers
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
	}	

	err := db.Debug().Create(&input).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request on header",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": input,
		"code" : http.StatusOK,
	})
}

