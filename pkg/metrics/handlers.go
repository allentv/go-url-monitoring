package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

// Register will add the above custom metrics to the Prom client
func Register() {
	prometheus.Register(sampleURLUp)
	prometheus.Register(sampleURLResponseTime)
}

// IncURLUp will update the value as up(1) or down(0) for an URL
func IncURLUp(url string, val bool) {
	var result float64

	if val {
		result = 1.0
	} else {
		result = 0.0
	}

	sampleURLUp.WithLabelValues(url).Set(result)
}

// IncURLResponseTime will add the response time for pinging an URL
func IncURLResponseTime(url string, ms uint64) {
	sampleURLResponseTime.WithLabelValues(url).Add(float64(ms))
}
