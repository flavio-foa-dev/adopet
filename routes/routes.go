package routes

import (
	"github.com/flavio-foa-dev/adopet/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		users := main.Group("users")
		{
			users.GET("/", controllers.ShowController)
		}
	}
	return router
}
