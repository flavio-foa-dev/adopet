package controllers

import (
	"fmt"
	"net/http"

	"github.com/flavio-foa-dev/adopet/database"
	"github.com/flavio-foa-dev/adopet/models"
	"github.com/gin-gonic/gin"
)

func ShowOKApi(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Message": "ok",
	})

}

func ShowPageHome(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", nil)
}

func CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	fmt.Println(&user, "created usuario")
	database.DB.Create(&user)
	c.JSON(http.StatusOK, &user)
}

func DeleteUser(c *gin.Context) {
	var user models.User
	id := c.Params.ByName("id")

	result := database.DB.Delete(&user, id)
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
			"message": "User not exist",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"user":    id,
	})

}

func ReativarUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	fmt.Println(id, "reativar pelo id")
	result := database.DB.Unscoped().Model(&user).Where("id = ? AND deleted_at IS NOT NULL", id).Update("deleted_at", nil)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, result.Error)
	}
	if result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, "user not found or already active")
	}
	c.JSON(http.StatusOK, gin.H{
		"ativado": true,
		"user":    user.Name,
		"id":      id,
	})
}
