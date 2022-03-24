package main

import (
	"github.com/allentv/go-url-monitoring/pkg/metrics"
	"github.com/allentv/go-url-monitoring/pkg/monitor"
	"github.com/allentv/go-url-monitoring/pkg/router"
	"github.com/allentv/go-url-monitoring/pkg/server"
)

func main() {
	metrics.Register()
	monitor.Start()

	srvConfig := server.NewDefaultConfig()
	srvConfig.UpdateHandler(router.New())

	server.Run(srvConfig)
}
