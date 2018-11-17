package main

import (
	"flag"
	"fmt"
	"mini-serve/route"
)

// ServerConfigs - store mini server config
type ServerConfigs struct {
	Host string
	Port int
}

func (serverConfigs *ServerConfigs) parseParameter() {
	flag.StringVar(&serverConfigs.Host, "HOST", "0.0.0.0", "Set binding ip")
	flag.IntVar(&serverConfigs.Port, "PORT", 9990, "Set binding port")
	flag.Parse()
}

func (serverConfigs ServerConfigs) toString() string {
	return fmt.Sprintf("%s:%d", serverConfigs.Host, serverConfigs.Port)
}

func main() {
	var serverConfigs ServerConfigs
	serverConfigs.parseParameter()
	miniServer := route.NewRoute()
	miniServer.Run(serverConfigs.toString())
}
