package router

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
)

type MonitoredRoutesResponse struct {
	URLs []string `json:"urls"`
}

// ListMonitoredRoutes will fetch the list of monitored routes
func ListMonitoredRoutes(w http.ResponseWriter, r *http.Request) {
	log.Ctx(r.Context()).Debug().Msg("In ListMonitoredRoutes handler")
	json.NewEncoder(w).Encode(MonitoredRoutesResponse{
		URLs: GetMonitoredURLs(),
	})
}
