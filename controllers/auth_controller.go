package controllers

import (
	"main/models"
	"main/services"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/now"
)

func GenerateToken(userId uint) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = userId
	claims["exp"] = now.EndOfMonth()
	claims["iat"] = time.Now().UTC().Unix()
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))

	return tokenString, err
}

type AuthController struct {
	authServices services.AuthServices
}

func NewAuthController(authServices services.AuthServices) *AuthController {
	return &AuthController{authServices: authServices}
}

func (a *AuthController) Login(ctx *gin.Context) {

	type User struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	var user User

	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{
			"status":  "Fail",
			"message": err.Error(),
		})
		return
	}

	result, err := a.authServices.Login(user.Email, user.Password)

	if err != nil {
		ctx.JSON(400, gin.H{
			"status":  "Fail",
			"message": err.Error(),
		})
		return
	}

	token, err := GenerateToken(result.ID)
	if err != nil {
		panic(err)
	}

	ctx.JSON(200, gin.H{
		"status": "ok",
		"user":   result,
		"token":  token,
	})
}

func (a *AuthController) SignUp(ctx *gin.Context) {
	var user models.User
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{
			"status":  "Fail",
			"message": err.Error(),
		})
		return
	}

	newUser, err := a.authServices.SignUp(user)

	if err != nil {
		panic(err.Error())
	}

	token, err := GenerateToken(newUser.ID)
	if err != nil {
		panic(err)
	}

	ctx.JSON(201, gin.H{
		"status": "Success",
		"user":   newUser,
		"token":  token,
	})
}

func (a *AuthController) ChangePassword(ctx *gin.Context) {
	type password struct {
		CurPassword string `json:"current_password" binding:"required"`
		NewPassword string `json:"new_password" binding:"required"`
	}

	var pass password

	if err := ctx.BindJSON(&pass); err != nil {
		ctx.JSON(400, gin.H{
			"status":  "Fail",
			"message": "Please enter current password and new password",
		})
		ctx.Abort()
		return
	}

	user, _ := ctx.Get("user")

	err := a.authServices.ChangePassword(user.(models.User).ID, pass.CurPassword, pass.NewPassword)
	if err != nil {
		ctx.JSON(400, gin.H{
			"status":  "Fail",
			"message": err.Error(),
		})
		ctx.Abort()
		return
	}

	token, err := GenerateToken(user.(models.User).ID)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Password change successfully",
		"token":   token,
	})

}
