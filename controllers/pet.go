package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/flavio-foa-dev/adopet/database"
	"github.com/flavio-foa-dev/adopet/models"
	"github.com/gin-gonic/gin"
)

func CreatePets(c *gin.Context) {
	var pet models.Pet
	var user models.User
	var shelter models.Shelter

	var data map[string]interface{}
	body, err := c.GetRawData()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error pegar bady": err.Error()})
		return
	}

	if err := json.Unmarshal(body, &data); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, ok1 := data["userID"].(string)
	shelterID, ok2 := data["shelterID"].(string)

	fmt.Println(userID, shelterID)
	if !ok1 || !ok2 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error map": "dados inválidos"})
		return
	}
	userid, _ := strconv.Atoi(userID)
	shelterid, _ := strconv.Atoi(userID)

	// faça algo com nome e idade

	database.DB.First(&user, userid)
	database.DB.First(&shelter, shelterid)

	pet.User = user
	pet.Shelter = shelter

	if err := c.ShouldBindJSON(&pet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error criacao": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Created", "pet": pet})

}
