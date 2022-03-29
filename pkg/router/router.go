package router

import (
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// New will configure the routes
func New() *mux.Router {
	router := mux.NewRouter()

	// For prometheus metrics
	router.Path("/metrics").Handler(promhttp.Handler())

	// Get list of all the routes
	router.HandleFunc("/routes", ListMonitoredRoutes)

	// Health check
	router.HandleFunc("/ping", HealthCheck)

	return router
}
