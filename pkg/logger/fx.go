package logger

import (
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

var Module = fx.Module("logger",
	fx.Provide(NewLogger),
)

var FxLogger = func(log *Logger) fxevent.Logger {
	return &fxevent.ZapLogger{Logger: log.zapLogger}
}
