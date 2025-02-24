package config

import (
	"go.uber.org/fx"
)

func WithConfig() fx.Option {
	return fx.Module("config",
		fx.Provide(NewEnvConfig),
	)
}
