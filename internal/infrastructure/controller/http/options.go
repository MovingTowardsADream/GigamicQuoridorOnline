package http

import (
	"time"
)

type Option func(*Server)

func Addr(addr string) Option {
	return func(s *Server) {
		s.httpServer.Addr = addr
	}
}

func ReadTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.httpServer.ReadTimeout = timeout
	}
}

func WriteTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.httpServer.WriteTimeout = timeout
	}
}

func ShutdownTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.shutdownTimeout = timeout
	}
}
