package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowController(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Message": "ok",
	})

}
