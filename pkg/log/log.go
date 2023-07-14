package log

import "go.uber.org/zap"

func Log() *zap.Logger {
	logger, _ := zap.NewProduction()

	return logger
}

func ErrorField(err error) zap.Field {
	return zap.String("error", err.Error())
}
