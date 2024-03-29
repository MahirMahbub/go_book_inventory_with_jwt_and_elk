package routes

import (
	"github.com/gin-gonic/gin"
	"go_practice/book/controllers"
	"go_practice/book/middlewares"
)

func AuthorRoute(v1 *gin.RouterGroup, c *controllers.Controller) {
	authorAdminGroup := v1.Group("/admin/authors").Use(middlewares.AdminAuth())
	{
		authorAdminGroup.POST("", c.CreateAdminAuthor)
		authorAdminGroup.GET(":id", c.FindAdminAuthor)
		authorAdminGroup.GET("", c.FindAdminAuthors)
		authorAdminGroup.PATCH(":id", c.UpdateAdminAuthor)
		authorAdminGroup.DELETE(":id", c.DeleteAdminAuthor)
	}
	authorGroup := v1.Group("/authors").Use(middlewares.Auth())
	{
		authorGroup.GET(":id", c.FindAuthor)
	}
}
