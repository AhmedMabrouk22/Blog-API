package routes

import (
	"main/config"
	"main/controllers"
	"main/middlewares"
	"main/repositories"
	"main/services"

	"github.com/gin-gonic/gin"
)

func CommentRoutes(r *gin.RouterGroup) {

	var db = config.DB
	commentRepository := repositories.NewCommentRepository(db)
	commentServices := services.NewCommentServices(commentRepository)
	commentController := controllers.NewCommentController(commentServices)
	r.GET("/:blogId", commentController.GetComments)

	r.Use(middlewares.Protect())
	r.POST("/:blogId", commentController.CreateComment)
	r.PATCH("/:commentId", commentController.UpdateComment)
	r.DELETE("/:commentId", commentController.DeleteComment)
}
