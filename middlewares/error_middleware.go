package middlewares

import (
	"os"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

func GlobalErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				if os.Getenv("ENV_MOD") == "development" {
					ctx.JSON(500, gin.H{
						"status":  "Error",
						"message": err,
						"stack":   string(debug.Stack()),
					})
				} else {
					ctx.JSON(500, gin.H{
						"status":  "Error",
						"message": err,
					})
				}
				ctx.Abort()
				return
			}
		}()

		ctx.Next()
	}
}
