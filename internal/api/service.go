package api

import (
	"github.com/mkramb/go-mockserver/pkg/logger"
)

type ApiService struct {
	log *logger.Logger
}

func NewApiService(log *logger.Logger) *ApiService {
	return &ApiService{
		log: log,
	}
}
