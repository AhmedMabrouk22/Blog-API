package main

import (
	"main/config"
	"main/middlewares"
	"main/models"
	"main/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	config.SetupDB()
	config.DB.AutoMigrate(&models.User{}, &models.Blog{}, &models.Topic{}, &models.Like{}, models.Comment{})
	server := gin.Default()
	server.Use(middlewares.GlobalErrorHandler())
	server.Static("/uploads", "./uploads")
	server.MaxMultipartMemory = 8 << 20
	router := server.Group("/api/v1")
	{
		userRouter := router.Group("/users")
		routes.UserRoutes(userRouter)
		topicRouter := router.Group("/topics")
		routes.TopicRoutes(topicRouter)
		blogRouter := router.Group("/blogs")
		routes.BlogRoutes(blogRouter)
		commentRouter := router.Group("/comments")
		routes.CommentRoutes(commentRouter)
		likeRouter := router.Group("/likes")
		routes.LikeRoutes(likeRouter)
	}

	server.Run()

}
