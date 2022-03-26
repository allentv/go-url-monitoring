package server

import (
	"net/http"
	"time"
)

// Config reprsents the configuration needed for the server
type Config struct {
	Domain         string
	Port           string
	Handler        http.Handler
	WriteTimeout   time.Duration
	ReadTimeout    time.Duration
	IdleTimeout    time.Duration
	DefaultTimeout time.Duration
}

// NewDefaultConfig will create a default server configuration
func NewDefaultConfig() *Config {
	return &Config{
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
func (s *Config) UpdatePort(port string) {
	s.Port = port
}

// UpdateHandler will update the HTTP handler
func (s *Config) UpdateHandler(h http.Handler) {
	s.Handler = h
}

// UpdateWriteTimeout will update the write timeout
func (s *Config) UpdateWriteTimeout(d time.Duration) {
	s.WriteTimeout = d
}

// UpdateWriteTimeout will update the read timeout
func (s *Config) UpdateReadTimeout(d time.Duration) {
	s.ReadTimeout = d
}

// UpdateWriteTimeout will update the idle timeout
func (s *Config) UpdateIdleTimeout(d time.Duration) {
	s.IdleTimeout = d
}

// UpdateDefaultTimeout will update the default timeout
func (s *Config) UpdateDefaultTimeout(d time.Duration) {
	s.DefaultTimeout = d
}
