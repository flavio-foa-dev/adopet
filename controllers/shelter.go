package controllers

import (
	"net/http"

	"github.com/flavio-foa-dev/adopet/database"
	"github.com/flavio-foa-dev/adopet/models"
	"github.com/gin-gonic/gin"
)

func CreateShelter(c *gin.Context) {
	var shelter models.Shelter

	if err := c.ShouldBindJSON(&shelter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	database.DB.Create(&shelter)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"shelter": shelter,
	})

}

func DeleteShelter(c *gin.Context) {
	var shelter models.Shelter
	var id = c.Params.ByName("id")

	result := database.DB.Delete(&shelter, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   result.Error.Error(),
		})
		return
	}
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNoContent, gin.H{
			"success": false,
			"message": "shelter not exist",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"user":    id,
	})

}

func ReativarShelter(c *gin.Context) {

}
