package api

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"

	"github.com/mkramb/go-mockserver/internal/models"
	"github.com/mkramb/go-mockserver/pkg/config"
	"github.com/mkramb/go-mockserver/pkg/logger"
)

type ApiServer struct {
	log    *logger.Logger
	cfg    *config.Config
	svcApi *ApiService
}

func NewApiServer(log *logger.Logger, cfg *config.Config, svcApi *ApiService) *ApiServer {
	return &ApiServer{
		log:    log,
		cfg:    cfg,
		svcApi: svcApi,
	}
}

func (s *ApiServer) CreateApiRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Recoverer)

	router.Post("/expectation", s.handleExpectations)
	router.Post("/verify", s.handleVerify)

	return router
}

func (s *ApiServer) handleExpectations(writer http.ResponseWriter, req *http.Request) {
	_, cancel := context.WithTimeout(context.Background(), s.cfg.DefaultTimeout)

	defer cancel()

	render.Status(req, http.StatusAccepted)
	render.JSON(writer, req, models.ApiResponse{
		Status: http.StatusAccepted,
	})
}

func (s *ApiServer) handleVerify(writer http.ResponseWriter, req *http.Request) {
	_, cancel := context.WithTimeout(context.Background(), s.cfg.DefaultTimeout)

	defer cancel()

	render.Status(req, http.StatusAccepted)
	render.JSON(writer, req, models.ApiResponse{
		Status: http.StatusAccepted,
	})
}
