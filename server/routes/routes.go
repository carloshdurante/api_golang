package routes

import (
	"github.com/carloshdurante/api_golang/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		books := main.Group("books")
		{
			books.GET("/:id", controllers.ShowBook)
			books.POST("/", controllers.CreateBook)
			books.DELETE("/:id", controllers.DeleteBook)
			books.PUT("/:id", controllers.UpdateBook)
			books.GET("/", controllers.ShowAllBooks)
		}
	}

	return router
}
