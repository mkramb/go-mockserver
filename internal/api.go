package internal

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/mkramb/go-mockserver/pkg/logger"
)

type ApiServer struct {
	log    *logger.Logger
	svcApi *ApiService
}

func NewApiServer(log *logger.Logger, svcApi *ApiService) *ApiServer {
	return &ApiServer{
		log:    log,
		svcApi: svcApi,
	}
}

func (s *ApiServer) CreateHttpRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Recoverer)

	return router
}
