package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Pong(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "Pong!",
	})
}

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		// tokenString := context.GetHeader(HEADER_KEY_AUTHORIZATION)
		fmt.Println("login access!")

		// if tokenString == "" {
		// 	context.JSON(http.StatusUnauthorized, gin.H{"error": NOT_CONTAIN_ACCESS_TOKEN})
		// 	context.Abort()
		// 	return
		// }
		// err := j.ValidateTokenJWT(tokenString)
		// if err != nil {
		// 	context.JSON(http.StatusUnauthorized, gin.H{"error2	": err.Error()})
		// 	context.Abort()
		// 	return
		// }
		context.Next()
	}
}
