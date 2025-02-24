package main

import (
	"go.uber.org/fx"

	"github.com/mkramb/go-mockserver/internal/api"
	"github.com/mkramb/go-mockserver/pkg/config"
	"github.com/mkramb/go-mockserver/pkg/logger"
	"github.com/mkramb/go-mockserver/pkg/server"
)

func Start(lc fx.Lifecycle, server *server.HttpServer, api *api.ApiServer) {
	go server.
		WithHttpRouter(api.CreateApiRouter()).
		StartHttp("go-mockserver")
}

func main() {
	fx.New(
		logger.Module,
		config.WithConfig(),
		server.NewHttpModule(),
		fx.WithLogger(logger.FxLogger),
		fx.Provide(api.NewApiService),
		fx.Provide(api.NewApiServer),
		fx.Invoke(Start),
	).Run()
}
