package middlewares

import (
	"main/config"
	"main/repositories"
	"main/services"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Protect() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := ctx.GetHeader("Authorization")

		if auth == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status":  "Fail",
				"message": "invalid token",
			})
			ctx.Abort()
			return
		}

		tokenString := strings.Split(auth, " ")[1]
		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status":  "Fail",
				"message": "invalid token",
			})
			ctx.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
		})

		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status":  "Fail",
				"message": "invalid token",
			})
			ctx.Abort()
			return
		}

		userRepository := repositories.NewUserRepository(config.DB)
		userServices := services.NewUserServices(userRepository)

		// Extract data from the token
		claims := token.Claims.(jwt.MapClaims)
		id := claims["id"]

		user, err := userServices.FindUserById(id)

		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"status":  "Fail",
				"message": "this user not found",
			})
			ctx.Abort()
			return
		}

		ctx.Set("user", user)

		ctx.Next()

	}
}
