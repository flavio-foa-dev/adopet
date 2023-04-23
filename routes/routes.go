package routes

import (
	"fmt"

	"github.com/flavio-foa-dev/adopet/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigureRoutes() {
	router := gin.Default()

	main := router.Group("api/v1")
	{
		main.GET("/ok", controllers.ShowOKApi)
	}
	{
		users := main.Group("users")
		{
			users.GET("/", controllers.ShowPageHome)
			users.POST("/", controllers.CreateUser)
			users.PUT("/reativar/:id", controllers.ReativarUser)
			users.DELETE("/:id", controllers.DeleteUser)
		}
	}
	{
		shelters := main.Group("shelters")
		{
			shelters.POST("/", controllers.CreateShelter)
		}
	}
	{
		pets := main.Group("pets")
		{
			pets.POST("/", controllers.CreatePets)
		}
	}

	fmt.Println("Starting Server...")
	router.Run(":8000")
}
