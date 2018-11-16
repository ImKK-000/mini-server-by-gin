package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ApiHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"ping":   "pong",
		"method": context.Request.Method,
	})
}
