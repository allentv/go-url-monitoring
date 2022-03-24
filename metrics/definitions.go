// All metrics definitions are made available here
package metrics

import "github.com/prometheus/client_golang/prometheus"

var sampleURLUp = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "sample_external_url_up",
		Help: "Whether the URL is up or not",
	},
	[]string{"url"},
)

var sampleURLResponseTime = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "sample_external_url_response_ms",
		Help: "Response time for the URL",
	},
	[]string{"url"},
)
