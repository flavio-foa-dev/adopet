package controllers

import (
	"net/http"

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
