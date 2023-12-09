package routes

import (
	"main/config"
	"main/controllers"
	"main/middlewares"
	"main/repositories"
	"main/services"

	"github.com/gin-gonic/gin"
)

func TopicRoutes(r *gin.RouterGroup) {

	var db = config.DB
	topicRepository := repositories.NewTopicRepository(db)
	topicServices := services.NewTopicServices(topicRepository)
	topicController := controllers.NewTopicController(topicServices)
	r.GET("/", topicController.FindAll)
	r.GET("/:id", topicController.FindTopic)

	r.Use(middlewares.Protect())
	r.POST("/", topicController.CreateTopic)

}
