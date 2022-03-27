package monitor

import (
	"time"

	"github.com/allentv/go-url-monitoring/pkg/router"
)

// Start will trigger the monitoring for all the available URLs
// by assigning a ticker per URL
func Start(urlsConfig []router.MonitoredURLConfig) int {
	for _, urlConfig := range urlsConfig {
		go func(config router.MonitoredURLConfig) {
			ticker := time.NewTicker(config.HealthCheckDuration)
			for {
				<-ticker.C
				WatchURL(config.Path)
			}
		}(urlConfig)
	}

	return len(urlsConfig)
}
