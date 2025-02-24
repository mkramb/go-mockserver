package main

import (
	"go.uber.org/fx"

	"github.com/mkramb/go-mockserver/internal"
	"github.com/mkramb/go-mockserver/pkg/logger"
	"github.com/mkramb/go-mockserver/pkg/server"
)

var cfg = NewEnvConfig()

func Start(lc fx.Lifecycle, server *server.HttpServer, api *internal.ApiServer) {
	go server.
		WithHttpRouter(api.CreateHttpRouter()).
		StartHttp("go-mockserver")
}

func main() {
	fx.New(
		logger.Module,
		server.NewHttpModule(cfg.HttpPort),
		fx.WithLogger(logger.FxLogger),
		fx.Provide(internal.NewApiService),
		fx.Provide(internal.NewApiServer),
		fx.Invoke(Start),
	).Run()
}
