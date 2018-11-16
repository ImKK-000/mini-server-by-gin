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

// ToString - convert host and port to new format
func (serverConfigs ServerConfigs) ToString() string {
	return fmt.Sprintf("%s:%d", serverConfigs.Host, serverConfigs.Port)
}

func main() {
	var serverConfigs ServerConfigs
	flag.StringVar(&serverConfigs.Host, "HOST", "0.0.0.0", "Set binding ip")
	flag.IntVar(&serverConfigs.Port, "PORT", 9990, "Set binding port")
	flag.Parse()

	miniServer := route.NewRoute()
	miniServer.Run(serverConfigs.ToString())
}
