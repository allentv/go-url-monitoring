package router_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/allentv/go-url-monitoring/pkg/router"
	"github.com/stretchr/testify/suite"
)

type HandlerSuite struct {
	suite.Suite
}

func TestHandlerSuite(t *testing.T) {
	suite.Run(t, new(HandlerSuite))
}

func (s *HandlerSuite) Test() {
	req, err := http.NewRequest("GET", "/routes", nil)
	s.Nil(err)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(router.ListMonitoredRoutes)
	handler.ServeHTTP(rr, req)
	s.Equal(http.StatusOK, rr.Code)
	s.Equal(
		fmt.Sprintln(`{"urls":["https://httpstat.us/503","https://httpstat.us/200"]}`),
		rr.Body.String(),
	)
}
