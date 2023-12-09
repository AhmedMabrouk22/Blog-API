package controllers

import (
	"main/models"
	"main/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type LikeController struct {
	LikeServices services.LikeServices
}

func NewLikeController(LikeServices services.LikeServices) *LikeController {
	return &LikeController{LikeServices: LikeServices}
}

func (c *LikeController) GetLikes(ctx *gin.Context) {
	blogId, err := strconv.ParseUint(ctx.Param("blogId"), 10, 32)
	if err != nil {
		panic(err)
	}

	likes, err := c.LikeServices.Get(uint(blogId))

	if err != nil {
		ctx.JSON(400, gin.H{
			"status":  "Fail",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(201, gin.H{
		"status": "Success",
		"result": len(likes),
		"likes":  likes,
	})

}

func (c *LikeController) AddLike(ctx *gin.Context) {
	blogId, err := strconv.ParseUint(ctx.Param("blogId"), 10, 32)
	if err != nil {
		panic(err)
	}

	data, _ := ctx.Get("user")
	curUser := data.(models.User)

	err = c.LikeServices.Add(uint(blogId), curUser.ID)

	if err != nil {
		ctx.JSON(400, gin.H{
			"status":  "Fail",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(201, gin.H{
		"status":  "Success",
		"message": "like added successfully",
	})
}

func (c *LikeController) DeleteLike(ctx *gin.Context) {
	blogId, err := strconv.ParseUint(ctx.Param("blogId"), 10, 32)
	if err != nil {
		panic(err)
	}

	data, _ := ctx.Get("user")
	curUser := data.(models.User)

	err = c.LikeServices.Delete(uint(blogId), curUser.ID)

	if err != nil {
		ctx.JSON(400, gin.H{
			"status":  "Fail",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"status":  "Success",
		"message": "like deleted successfully",
	})
}
