package router

import (
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// New will configure the routes
func New() *mux.Router {
	router := mux.NewRouter()

	router.Path("/metrics").Handler(promhttp.Handler())

	router.HandleFunc("/routes", ListMonitoredRoutes)

	return router
}
