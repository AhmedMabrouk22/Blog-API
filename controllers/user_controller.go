package controllers

import (
	"fmt"
	"main/models"
	"main/services"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type userController struct {
	userServices services.UserServices
}

func NewUserController(userServices services.UserServices) *userController {
	return &userController{userServices: userServices}
}

func (u *userController) GetMe(ctx *gin.Context) {
	value, exists := ctx.Get("user")

	if !exists {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  "Fail",
			"message": "User not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "Success",
		"data":   value,
	})
}

func (u *userController) UpdateMe(ctx *gin.Context) {
	type User struct {
		Name  string `json:"name" form:"name"`
		Image string
	}

	var newUser User
	if err := ctx.ShouldBind(&newUser); err != nil {
		ctx.JSON(400, gin.H{
			"status":  "Fail",
			"message": err.Error(),
		})
		ctx.Abort()
		return
	}

	curUser, _ := ctx.Get("user")
	updatedUser := curUser.(models.User)
	if newUser.Name != "" {
		updatedUser.Name = newUser.Name
	}

	file, _ := ctx.FormFile("image")

	if file != nil {
		fileType := strings.Split(file.Header.Values("Content-Type")[0], "/")[1]

		data, _ := ctx.Get("user")
		curUser := data.(models.User)

		imageName := fmt.Sprintf("user-%v-%v.%v", curUser.ID, time.Now().Unix(), fileType)
		path := fmt.Sprintf("uploads/users/%v", imageName)
		ctx.SaveUploadedFile(file, path)
		updatedUser.Image = imageName
	}

	res, err := u.userServices.Update(curUser.(models.User).ID, updatedUser)

	if err != nil {
		ctx.JSON(400, gin.H{
			"status":  "Fail",
			"message": err.Error(),
		})
		ctx.Abort()
		return
	}

	ctx.JSON(200, gin.H{
		"status": "Success",
		"user":   res,
	})
}

func (u *userController) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := u.userServices.GetUser(id)

	if err != nil {
		ctx.JSON(404, gin.H{
			"status":  "Fail",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"status": "Success",
		"user":   user,
	})
}
