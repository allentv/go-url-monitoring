package server

import (
	"net/http"
	"time"
)

// ServerConfig reprsents the configuration needed for the server
type ServerConfig struct {
	Domain         string
	Port           string
	Handler        http.Handler
	WriteTimeout   time.Duration
	ReadTimeout    time.Duration
	IdleTimeout    time.Duration
	DefaultTimeout time.Duration
}

// NewDefaultConfig will create a default server configuration
func NewDefaultConfig() *ServerConfig {
	return &ServerConfig{
		Domain:         "0.0.0.0",
		Port:           "9000",
		Handler:        nil,
		WriteTimeout:   time.Second * 15,
		ReadTimeout:    time.Second * 15,
		IdleTimeout:    time.Second * 60,
		DefaultTimeout: time.Second * 15,
	}
}

// UpdatePort will update the port
func (s *ServerConfig) UpdatePort(port string) {
	s.Port = port
}

// UpdateHandler will update the HTTP handler
func (s *ServerConfig) UpdateHandler(h http.Handler) {
	s.Handler = h
}

// UpdateWriteTimeout will update the write timeout
func (s *ServerConfig) UpdateWriteTimeout(d time.Duration) {
	s.WriteTimeout = d
}

// UpdateWriteTimeout will update the read timeout
func (s *ServerConfig) UpdateReadTimeout(d time.Duration) {
	s.ReadTimeout = d
}

// UpdateWriteTimeout will update the idle timeout
func (s *ServerConfig) UpdateIdleTimeout(d time.Duration) {
	s.IdleTimeout = d
}

// UpdateDefaultTimeout will update the default timeout
func (s *ServerConfig) UpdateDefaultTimeout(d time.Duration) {
	s.DefaultTimeout = d
}
