package monitor_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/allentv/go-url-monitoring/pkg/monitor"
	"github.com/allentv/go-url-monitoring/pkg/router"
	"github.com/stretchr/testify/suite"
)

type MonitorSuite struct {
	suite.Suite
}

func TestMonitorSuite(t *testing.T) {
	suite.Run(t, new(MonitorSuite))
}

func (s *MonitorSuite) TestStart() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("."))
	}))
	defer ts.Close()

	urlConfig := []router.MonitoredURLConfig{
		{
			Path:                ts.URL + "/ping",
			HealthCheckDuration: 1 * time.Millisecond,
		},
	}

	s.Equal(1, monitor.Start(urlConfig))
}
