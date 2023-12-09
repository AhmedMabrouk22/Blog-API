package middlewares

import (
	"os"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

func GlobalErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		func() {
			if err := recover(); err != nil {
				if os.Getenv("ENV_MOD") == "development" {
					ctx.JSON(500, gin.H{
						"status":  "Error",
						"message": err,
						"stack":   string(debug.Stack()),
					})
				}
				ctx.Abort()
				return
			}
		}()

		ctx.Next()
	}
}