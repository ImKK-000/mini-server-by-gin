package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler - use when call path /say
func Handler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"ping":   "pong",
		"method": context.Request.Method,
	})
}
