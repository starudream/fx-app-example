package log

import (
	"github.com/starudream/go-lib/v2/config"

	"github.com/starudream/go-lib/v2/internal/log"
)

var defaultLogger *SugaredLogger

func NewFromKoanf(k *config.Koanf) (*Logger, *SugaredLogger) {
	cfg := Config{}
	err := k.Unmarshal("log", &cfg)
	if err != nil {
		log.X.Named("log").Fatal("parse config error: %v", err)
	}

	if config.Debug(k) {
		cfg.Console.Level = "debug"
	}

	logger := New(cfg)

	ReplaceGlobals(logger)
	RedirectStdLog(logger)

	defaultLogger = logger.Sugar()

	return logger, defaultLogger
}

func L() *SugaredLogger {
	return defaultLogger
}
