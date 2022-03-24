package main

import (
	"github.com/allentv/go-url-monitoring/metrics"
	"github.com/allentv/go-url-monitoring/monitor"
	"github.com/allentv/go-url-monitoring/router"
	"github.com/allentv/go-url-monitoring/server"
)

func main() {
	metrics.Register()
	monitor.Start()

	srvConfig := server.NewDefaultConfig()
	srvConfig.UpdateHandler(router.New())

	server.Run(srvConfig)
}
