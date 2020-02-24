package log

import (
	"go.uber.org/zap"
)

var LOG *zap.SugaredLogger

func init() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	LOG = logger.Sugar()
}
