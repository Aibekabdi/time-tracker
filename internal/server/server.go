package server

import (
	"context"
	"net/http"
	"time-tracker/internal/config"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(cfg config.HTTPServer, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:         cfg.Address,
		Handler:      handler,
		ReadTimeout:  cfg.Timeout,
		WriteTimeout: cfg.Timeout,
		IdleTimeout:  cfg.Timeout,
	}
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
