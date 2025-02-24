package logger

import "go.uber.org/zap"

func AsError(val error) zap.Field {
	return zap.Error(val)
}
