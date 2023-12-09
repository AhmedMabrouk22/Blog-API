package routes

import (
	"main/config"
	"main/controllers"
	"main/middlewares"
	"main/repositories"
	"main/services"

	"github.com/gin-gonic/gin"
)

func LikeRoutes(r *gin.RouterGroup) {

	var db = config.DB
	likeRepository := repositories.NewLikeRepository(db)
	likeServices := services.NewLikeServices(likeRepository)
	likeController := controllers.NewLikeController(likeServices)
	r.GET("/:blogId", likeController.GetLikes)
	r.Use(middlewares.Protect())
	r.POST("/add/:blogId", likeController.AddLike)
	r.DELETE("/delete/:blogId", likeController.DeleteLike)
}
