package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/isaqueveras/rest-api-gin-postgres-gorm/controllers"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		books := main.Group("books")
		{
			books.GET("", controllers.ShowBooks)
			books.GET("/:id", controllers.ShowBook)
			books.POST("", controllers.CreateBook)
			books.PUT("/", controllers.UpdateBook)
			books.DELETE("/:id", controllers.DeleteBook)
		}
	}

	return router
}
