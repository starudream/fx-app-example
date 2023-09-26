package log

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var X = zap.New(zapcore.NewCore(zapcore.NewConsoleEncoder(NewEncoderConfig(true)), os.Stdout, zapcore.InfoLevel)).Sugar()
