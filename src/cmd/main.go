package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ServerConfigs store mini server config
type ServerConfigs struct {
	Host string
	Port int
}

// ToString convert host and port to new format
func (serverConfigs ServerConfigs) ToString() string {
	return fmt.Sprintf("%s:%d", serverConfigs.Host, serverConfigs.Port)
}

func apiHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"ping":   "pong",
		"method": context.Request.Method,
	})
}

func main() {
	var serverConfigs ServerConfigs
	flag.StringVar(&serverConfigs.Host, "HOST", "0.0.0.0", "Set binding ip")
	flag.IntVar(&serverConfigs.Port, "PORT", 9990, "Set binding port")
	flag.Parse()

	route := gin.Default()
	route.GET("say", apiHandler)
	route.POST("say", apiHandler)
	route.PUT("say", apiHandler)
	route.DELETE("say", apiHandler)
	route.Run(serverConfigs.ToString())
}
