package routes

import (
	"main/config"
	"main/controllers"
	"main/middlewares"
	"main/repositories"
	"main/services"

	"github.com/gin-gonic/gin"
)

func BlogRoutes(r *gin.RouterGroup) {

	var db = config.DB
	blogRepository := repositories.NewBlogRepository(db)
	blogServices := services.NewBlogServices(blogRepository)
	blogController := controllers.NewBlogController(blogServices)

	r.GET("/", blogController.FindAll)
	r.GET("/:id", blogController.FindBlog)

	r.Use(middlewares.Protect())
	r.POST("/", blogController.CreateBlog)
	r.PUT("/:id", blogController.UpdateBlog)
	r.DELETE("/:id", blogController.DeleteBlog)

}
