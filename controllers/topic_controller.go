package controllers

import (
	"main/models"
	"main/services"

	"github.com/gin-gonic/gin"
)

type TopicController struct {
	topicServices services.TopicServices
}

func NewTopicController(topicServices services.TopicServices) *TopicController {
	return &TopicController{topicServices: topicServices}
}

func (t *TopicController) CreateTopic(ctx *gin.Context) {
	var topic models.Topic

	if err := ctx.ShouldBind(&topic); err != nil {
		ctx.JSON(400, gin.H{
			"status":  "Fail",
			"message": err.Error(),
		})
		return
	}

	newTopic := t.topicServices.Create(topic)
	ctx.JSON(201, gin.H{
		"status": "Success",
		"topic":  newTopic,
	})
}

func (t *TopicController) FindAll(ctx *gin.Context) {
	topics := t.topicServices.FindAll()

	ctx.JSON(200, gin.H{
		"status": "Success",
		"result": len(topics),
		"topics": topics,
	})
}

func (t *TopicController) FindTopic(ctx *gin.Context) {
	id := ctx.Param("id")
	topic, err := t.topicServices.Find(id)

	if err != nil {
		ctx.JSON(404, gin.H{
			"status":  "Fail",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"status":      "Success",
		"Topic Name":  topic.Name,
		"Blog Number": len(topic.Blogs),
		"Blogs":       topic.Blogs,
	})

}
