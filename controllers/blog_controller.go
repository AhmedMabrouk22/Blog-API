package controllers

import (
	"fmt"
	"main/models"
	"main/services"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type BlogController struct {
	blogServices services.BlogServices
}

func NewBlogController(blogServices services.BlogServices) *BlogController {
	return &BlogController{blogServices: blogServices}
}

func getBlog(ctx *gin.Context, blog models.BlogRequest) models.BlogRequest {
	data, _ := ctx.Get("user")
	curUser := data.(models.User)
	blog.AuthorID = curUser.ID
	file, _ := ctx.FormFile("cover")
	if file != nil {
		fileType := strings.Split(file.Header.Values("Content-Type")[0], "/")[1]
		imageName := fmt.Sprintf("blog-%v-%v.%v", curUser.ID, time.Now().Unix(), fileType)
		path := fmt.Sprintf("uploads/blogs/%v", imageName)
		ctx.SaveUploadedFile(file, path)
		blog.ImageCover = imageName
	}
	return blog
}

func (controller *BlogController) CreateBlog(ctx *gin.Context) {
	var blog models.BlogRequest

	if err := ctx.ShouldBind(&blog); err != nil {
		ctx.JSON(400, gin.H{
			"status":  "Fail",
			"message": err.Error(),
		})
		return
	}
	blog = getBlog(ctx, blog)
	newBlog := controller.blogServices.Create(blog)
	ctx.JSON(201, gin.H{
		"status": "Success",
		"blog":   newBlog,
	})
}

func (controller *BlogController) UpdateBlog(ctx *gin.Context) {
	id := ctx.Param("id")
	var blog models.BlogRequest

	if err := ctx.ShouldBind(&blog); err != nil {
		ctx.JSON(400, gin.H{
			"status":  "Fail",
			"message": err.Error(),
		})
		return
	}
	blog = getBlog(ctx, blog)
	err := controller.blogServices.Update(id, blog)
	if err != nil {
		ctx.JSON(404, gin.H{
			"status":  "Fail",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"status":  "Success",
		"message": "blog update successfully",
	})
}

func (controller *BlogController) DeleteBlog(ctx *gin.Context) {
	id := ctx.Param("id")
	err := controller.blogServices.Delete(id)
	if err != nil {
		ctx.JSON(404, gin.H{
			"status":  "Fail",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"status":  "Success",
		"message": "Blog deleted successfully",
	})
}

func (controller *BlogController) FindBlog(ctx *gin.Context) {
	id := ctx.Param("id")
	blog, err := controller.blogServices.Find(id)
	if err != nil {
		ctx.JSON(404, gin.H{
			"status":  "Fail",
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"status": "Success",
		"blog":   blog,
	})
}

func (controllers *BlogController) FindAll(ctx *gin.Context) {

	blogs, err := controllers.blogServices.FindAll()
	if err != nil {
		if err != nil {
			ctx.JSON(400, gin.H{
				"status":  "Fail",
				"message": err.Error(),
			})
			return
		}
	}

	ctx.JSON(200, gin.H{
		"status": "Success",
		"result": len(blogs),
		"blogs":  blogs,
	})
}
