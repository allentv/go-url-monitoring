package router

import "time"

// MonitoredURLConfig represents the health check configuration per monitored URL
type MonitoredURLConfig struct {
	Path                string
	HealthCheckDuration time.Duration
}

const defaultHealthCheckDuration = 5 * time.Second

// GetMonitoredURLsConfig will return the configuration per URL
func GetMonitoredURLsConfig() []MonitoredURLConfig {
	result := []MonitoredURLConfig{}

	for _, url := range GetMonitoredURLs() {
		result = append(result, MonitoredURLConfig{
			Path:                url,
			HealthCheckDuration: defaultHealthCheckDuration, // this would be different per URL
		})
	}

	return result
}

// GetMonitoredURLs will return the list of URLs to be monitored
func GetMonitoredURLs() []string {
	// The list can be loaded from some persistent data store
	// or by calling an endpoint which in turn will add the new
	// route and the the time interval for the healthcheck to take
	// place or it can be loaded from another endpoint

	return []string{
		"https://httpstat.us/503",
		"https://httpstat.us/200",
	}
}
