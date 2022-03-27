package main

import (
	"github.com/allentv/go-url-monitoring/pkg/metrics"
	"github.com/allentv/go-url-monitoring/pkg/monitor"
	"github.com/allentv/go-url-monitoring/pkg/router"
	"github.com/allentv/go-url-monitoring/pkg/server"
)

func main() {
	// Register available custom metrics
	metrics.Register()

	// Start a monitoring process for each of the existing URLs
	monitor.Start(router.GetMonitoredURLsConfig())

	srvConfig := server.NewDefaultConfig()
	srvConfig.UpdateHandler(router.New())

	server.Run(srvConfig)
}
