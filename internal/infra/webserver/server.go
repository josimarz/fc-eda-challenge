package webserver

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

type Handler interface {
	GetMethod() string
	GetPattern() string
	GetHandlerFunc() http.HandlerFunc
}

type Server struct {
	Port     string
	Router   chi.Router
	Handlers []Handler
}

func NewServer(port string) *Server {
	return &Server{
		Port:     port,
		Router:   chi.NewRouter(),
		Handlers: []Handler{},
	}
}

func (s *Server) AddHandler(handler Handler) {
	s.Handlers = append(s.Handlers, handler)
}

func (s *Server) Start() error {
	s.Router.Use(middleware.Logger)
	for _, handler := range s.Handlers {
		s.Router.MethodFunc(
			handler.GetMethod(),
			handler.GetPattern(),
			handler.GetHandlerFunc(),
		)
	}
	return http.ListenAndServe(s.Port, s.Router)
}
