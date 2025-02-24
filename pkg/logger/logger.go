package logger

import (
	"go.uber.org/zap"
)

type Logger struct {
	*zap.SugaredLogger
	zapLogger *zap.Logger
}

func NewLogger() *Logger {
	zapLogger, err := zap.NewProduction()

	if err != nil {
		panic("Error initializing logger")
	}

	return &Logger{
		SugaredLogger: zapLogger.Sugar(),
		zapLogger:     zapLogger,
	}
}
