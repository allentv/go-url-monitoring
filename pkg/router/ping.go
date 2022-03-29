package router

import "net/http"

// HealthCheck is the health check handler
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("."))
}
