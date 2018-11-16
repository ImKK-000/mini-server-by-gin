package route

import (
	"mini-serve/api"

	"github.com/gin-gonic/gin"
)

func NewRoute() *gin.Engine {
	route := gin.Default()
	route.GET("say", api.ApiHandler)
	route.POST("say", api.ApiHandler)
	route.PUT("say", api.ApiHandler)
	route.DELETE("say", api.ApiHandler)
	return route
}
