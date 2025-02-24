package server

import (
	"context"

	"github.com/mkramb/go-mockserver/pkg/logger"
	"go.uber.org/fx"
)

func NewHttpModule(httpPort int) fx.Option {
	return fx.Module("transport",
		fx.Provide(func(lc fx.Lifecycle, log *logger.Logger) *HttpServer {
			server := NewHttpServer(log, httpPort)

			lc.Append(fx.Hook{
				OnStop: func(ctx context.Context) error {
					log.Infow("Closing http server")

					if err := server.http.Shutdown(ctx); err != nil {
						log.Error("Error stopping http server")
					}

					return nil
				},
			})

			return server
		}),
	)
}
