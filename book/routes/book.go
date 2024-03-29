package routes

import (
	"github.com/gin-gonic/gin"
	"go_practice/book/controllers"
	"go_practice/book/middlewares"
)

func BookRoute(v1 *gin.RouterGroup, c *controllers.Controller) {
	bookGroup := v1.Group("/books").Use(middlewares.Auth())
	{
		bookGroup.GET(":id", c.FindBook)
		bookGroup.GET("", c.FindUserBooks)
		bookGroup.POST("", c.CreateBook)
		bookGroup.DELETE(":id", c.DeleteBook)
		bookGroup.PATCH(":id", c.UpdateBook)
	}
	bookAdminGroup := v1.Group("/admin/books").Use(middlewares.AdminAuth())
	{
		bookAdminGroup.GET(":id", c.FindAdminBook)
		bookAdminGroup.GET("", c.FindAdminBooks)
		bookAdminGroup.POST("", c.CreateAdminBook)
		bookAdminGroup.DELETE(":id", c.DeleteAdminBook)
		bookAdminGroup.PATCH(":id", c.UpdateAdminBook)
	}
}
