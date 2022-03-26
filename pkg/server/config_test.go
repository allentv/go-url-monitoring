package server_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/allentv/go-url-monitoring/pkg/server"
	"github.com/stretchr/testify/suite"
)

type ConfigSuite struct {
	suite.Suite
	sc *server.Config
}

func TestConfigSuite(t *testing.T) {
	suite.Run(t, new(ConfigSuite))
}

func (s *ConfigSuite) SetupTest() {
	s.sc = server.NewDefaultConfig()
}

func (s *ConfigSuite) TestNewDefaultConfig() {
	s.Equal("0.0.0.0", s.sc.Domain)
	s.Equal("9000", s.sc.Port)
	s.Nil(s.sc.Handler)
	s.Equal(15*time.Second, s.sc.WriteTimeout)
	s.Equal(15*time.Second, s.sc.ReadTimeout)
	s.Equal(60*time.Second, s.sc.IdleTimeout)
	s.Equal(15*time.Second, s.sc.DefaultTimeout)
}

func (s *ConfigSuite) TestUpdatePort() {
	s.sc.UpdatePort("3000")
	s.Equal("3000", s.sc.Port)
}

func (s *ConfigSuite) TestUpdateHandler() {
	s.sc.UpdateHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	s.NotNil(s.sc.Handler)
}

func (s *ConfigSuite) TestUpdateWriteTimeout() {
	s.sc.UpdateWriteTimeout(1 * time.Second)
	s.Equal(1*time.Second, s.sc.WriteTimeout)
}

func (s *ConfigSuite) TestUpdateReadTimeout() {
	s.sc.UpdateReadTimeout(1 * time.Second)
	s.Equal(1*time.Second, s.sc.ReadTimeout)
}

func (s *ConfigSuite) TestUpdateIdleTimeout() {
	s.sc.UpdateIdleTimeout(1 * time.Second)
	s.Equal(1*time.Second, s.sc.IdleTimeout)
}

func (s *ConfigSuite) TestUpdateDefaultTimeout() {
	s.sc.UpdateDefaultTimeout(1 * time.Second)
	s.Equal(1*time.Second, s.sc.DefaultTimeout)
}
