package router

import (
	"encoding/json"
	"net/http"
)

type MonitoredRoutesResponse struct {
	URLs []string `json:"urls"`
}

// ListMonitoredRoutes will fetch the list of monitored routes
func ListMonitoredRoutes(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(MonitoredRoutesResponse{
		URLs: GetMonitoredURLs(),
	})
}
