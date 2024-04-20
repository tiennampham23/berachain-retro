package log

import (
	"go.uber.org/zap"
)

var (
	l *zap.SugaredLogger
)

func init() {
	logger, _ := zap.NewProduction()
	l = logger.Sugar()
}

func Logger() *zap.SugaredLogger {
	return l
}
