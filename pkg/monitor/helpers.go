package monitor

import (
	"net/http"
	"time"

	"github.com/allentv/go-url-monitoring/pkg/metrics"
)

// WatchURL will monitor the uptime for an URL
func WatchURL(url string) {
	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		metrics.IncURLUp(url, false)
		return
	}

	defer resp.Body.Close()

	// This is a naive implementation
	// Using HTTP Roundtripper would provide hooks for
	// tracking time for specific areas in establishing the connection
	elapsed := time.Since(start).Milliseconds()

	metrics.IncURLUp(url, resp.StatusCode == http.StatusOK)
	metrics.IncURLResponseTime(url, uint64(elapsed))
}
