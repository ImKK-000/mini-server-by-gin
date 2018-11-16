package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	var host string
	flag.StringVar(&host, "HOST", "0.0.0.0", "Set binding ip (default: 0.0.0.0)")
	var port int
	flag.IntVar(&port, "PORT", 9990, "Set binding port (default: 9990)")
	flag.Parse()

	route := gin.Default()
	route.GET("say", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"ping": "pong",
		})
	})
	route.Run(fmt.Sprintf("%s:%d", host, port))
}
