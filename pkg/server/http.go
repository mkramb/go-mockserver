package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/hellofresh/health-go/v5"
	"github.com/mkramb/go-mockserver/pkg/logger"
)

type HttpServer struct {
	apiRouter *chi.Mux
	http      *http.Server
	log       *logger.Logger
}

func NewHttpServer(log *logger.Logger, port int) *HttpServer {
	server := http.Server{
		Addr:        fmt.Sprintf(":%d", port),
		IdleTimeout: time.Minute,
	}

	return &HttpServer{
		apiRouter: nil,
		http:      &server,
		log:       log,
	}
}

func (s *HttpServer) WithHttpRouter(apiRouter *chi.Mux) *HttpServer {
	s.apiRouter = apiRouter

	return s
}

func (s *HttpServer) StartHttp(serviceName string) {
	s.log.Infow("Starting HTTP server", "address", s.http.Addr)
	s.http.Handler = s.CreateHttpRouter(health.Component{
		Name:    serviceName,
		Version: "1.0.0",
	})

	go func() {
		err := s.http.ListenAndServe()

		if err != nil && err != http.ErrServerClosed {
			s.log.Infow("Error starting http server", logger.AsError(err))
		}
	}()
}

func (s *HttpServer) CreateHttpRouter(serviceInfo health.Component) *chi.Mux {
	router := chi.NewRouter()
	status, err := health.New(health.WithComponent(serviceInfo))

	if err != nil {
		s.log.Errorw("Error initializing http server", logger.AsError(err))
		panic("Error initializing http server")
	}

	router.Mount("/health", status.Handler())

	if s.apiRouter != nil {
		router.Mount("/mockserver", s.apiRouter)
	}

	return router
}
