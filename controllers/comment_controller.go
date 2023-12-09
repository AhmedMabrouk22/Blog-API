package controllers

import (
	"main/models"
	"main/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentController struct {
	commentServices services.CommentServices
}

func NewCommentController(commentServices services.CommentServices) *CommentController {
	return &CommentController{commentServices: commentServices}
}

func (c *CommentController) CreateComment(ctx *gin.Context) {
	blogId, err := strconv.ParseUint(ctx.Param("blogId"), 10, 32)
	if err != nil {
		panic(err)
	}
	type commentRequest struct {
		Content string `json:"content" binding:"required" form:"content"`
	}
	var commentRec commentRequest

	if err := ctx.ShouldBind(&commentRec); err != nil {
		ctx.JSON(400, gin.H{
			"status":  "Fail",
			"message": err.Error(),
		})
		return
	}

	data, _ := ctx.Get("user")
	curUser := data.(models.User)

	var comment models.Comment
	comment.Content = commentRec.Content
	comment.BlogID = uint(blogId)
	comment.UserID = curUser.ID
	comment.User = curUser
	newComment, err := c.commentServices.Create(comment)
	if err != nil {
		ctx.JSON(400, gin.H{
			"status":  "Fail",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(201, gin.H{
		"status":  "Success",
		"comment": newComment,
	})
}

func (c *CommentController) UpdateComment(ctx *gin.Context) {
	commentId := ctx.Param("commentId")

	type commentRequest struct {
		Content string `json:"content" binding:"required" form:"content"`
	}
	var commentRec commentRequest

	if err := ctx.ShouldBind(&commentRec); err != nil {
		ctx.JSON(400, gin.H{
			"status":  "Fail",
			"message": err.Error(),
		})
		return
	}

	var comment models.Comment
	comment.Content = commentRec.Content
	data, _ := ctx.Get("user")
	curUser := data.(models.User)
	err := c.commentServices.Update(commentId, curUser.ID, comment)

	if err != nil {
		ctx.JSON(400, gin.H{
			"status":  "Fail",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"status":  "Success",
		"message": "comment updated successfully",
	})
}

func (c *CommentController) DeleteComment(ctx *gin.Context) {
	commentId := ctx.Param("commentId")
	data, _ := ctx.Get("user")
	curUser := data.(models.User)
	err := c.commentServices.Delete(commentId, curUser.ID)

	if err != nil {
		ctx.JSON(400, gin.H{
			"status":  "Fail",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"status":  "Success",
		"message": "comment deleted successfully",
	})
}

func (c *CommentController) GetComments(ctx *gin.Context) {
	blogId := ctx.Param("blogId")
	comments, err := c.commentServices.Get(blogId)
	if err != nil {
		ctx.JSON(400, gin.H{
			"status":  "Fail",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"status":   "Success",
		"comments": comments,
	})
}
