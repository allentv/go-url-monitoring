package router_test

import (
	"testing"

	"github.com/allentv/go-url-monitoring/pkg/router"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/suite"
)

type RouterSuite struct {
	suite.Suite
}

func TestRouterSuite(t *testing.T) {
	suite.Run(t, new(RouterSuite))
}

func (s *RouterSuite) TestNew() {
	r := router.New()
	r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		pathTemplate, err := route.GetPathTemplate()
		if err == nil {
			s.Contains(
				[]string{
					"/metrics",
					"/routes",
					"/ping",
				},
				pathTemplate,
			)
		}

		return nil
	})
}
