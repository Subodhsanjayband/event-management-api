package middlewares

import (
	"net/http"

	"github.com/Subodhsanjayband/event_manager/utils"
	"github.com/gin-gonic/gin"
)

func authenticate(context *gin.Context) {
	token := context.Request.Header.Get("authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Not Authorized",
		})
	}

	id, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Not authorised",
		})
		return
	}
	context.Next()
}
