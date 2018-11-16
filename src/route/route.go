package route

import (
	"mini-serve/api"

	"github.com/gin-gonic/gin"
)

// NewRoute - create new server with route
func NewRoute() *gin.Engine {
	route := gin.Default()
	route.GET("say", api.Handler)
	route.POST("say", api.Handler)
	route.PUT("say", api.Handler)
	route.DELETE("say", api.Handler)
	return route
}
