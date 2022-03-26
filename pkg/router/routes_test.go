package router_test

import (
	"testing"
	"time"

	"github.com/allentv/go-url-monitoring/pkg/router"
	"github.com/stretchr/testify/suite"
)

type RoutesSuite struct {
	suite.Suite
}

func TestRoutesSuite(t *testing.T) {
	suite.Run(t, new(RoutesSuite))
}

func (s *RoutesSuite) TestGetMonitoredURLsConfig() {
	for _, uConfig := range router.GetMonitoredURLsConfig() {
		s.NotEmpty(uConfig.Path)
		s.Equal(5*time.Second, uConfig.HealthCheckDuration)
	}
}

func (s *RoutesSuite) TestGetMonitoredURLs() {
	// Check if there is atleast one route available for monitoring by default
	s.True(len(router.GetMonitoredURLs()) > 0)
}
