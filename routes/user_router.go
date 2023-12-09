package routes

import (
	"main/config"
	"main/controllers"
	"main/middlewares"
	"main/repositories"
	"main/services"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup) {

	var db = config.DB
	userRepository := repositories.NewUserRepository(db)
	authServices := services.NewAuthServices(userRepository)
	authController := controllers.NewAuthController(authServices)
	r.POST("/login", authController.Login)
	r.POST("/signup", authController.SignUp)

	userServices := services.NewUserServices(userRepository)
	userController := controllers.NewUserController(userServices)

	r.GET("/:id", userController.GetUser)
	r.Use(middlewares.Protect())
	r.GET("/me", userController.GetMe)
	r.PATCH("/changePassword", authController.ChangePassword)
	r.PATCH("/update", userController.UpdateMe)

}
