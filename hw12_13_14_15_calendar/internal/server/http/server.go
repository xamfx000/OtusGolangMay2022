package internalhttp

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	logger Logger
	Host   string
	Port   int
}

type Logger interface {
	Info(msg string)
	Error(msg string)
	Warn(msg string)
	Debug(msg string)
}

type Application interface { // TODO
}

func NewServer(logger Logger, app Application) *Server {
	return &Server{logger: logger}
}

func (s *Server) Start(ctx context.Context) error {
	r := chi.NewRouter()
	r.Use(loggingMiddleware)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {})
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", s.Host, s.Port), r)
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	// TODO
	return nil
}

// TODO
