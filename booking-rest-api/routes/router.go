package routes

import (
	"webapi-with-go/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {

	//Creating endPoints
	main := router.Group("api")
	{
		books := main.Group("books")
		{
			books.GET("/", controllers.ShowBook)
		}
	}

	return router
}
