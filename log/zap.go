package log

import (
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/starudream/go-lib/v2/osx"

	"github.com/starudream/go-lib/v2/internal/log"
)

type (
	Logger        = zap.Logger
	SugaredLogger = zap.SugaredLogger
)

var (
	ReplaceGlobals = zap.ReplaceGlobals
	RedirectStdLog = zap.RedirectStdLog

	WithCaller    = zap.WithCaller
	AddCallerSkip = zap.AddCallerSkip

	Fields = zap.Fields
	Any    = zap.Any
)

type Config struct {
	Console ConfigConsole `yaml:"console" koanf:"console"`
	File    ConfigFile    `yaml:"file" koanf:"file"`
}

type ConfigConsole struct {
	Level   string `yaml:"level" koanf:"level"`
	NoColor bool   `yaml:"no_color" koanf:"no_color"`
}

type ConfigFile struct {
	Level string `yaml:"level" koanf:"level"`
	Color bool   `yaml:"color" koanf:"color"`

	Filename   string `yaml:"filename" koanf:"filename"`
	MaxSize    int    `yaml:"max_size" koanf:"max_size"`
	MaxAge     int    `yaml:"max_age" koanf:"max_age"`
	MaxBackups int    `yaml:"max_backups" koanf:"max_backups"`
	Compress   bool   `yaml:"compress" koanf:"compress"`
}

func New(cfg Config) *Logger {
	var cores []zapcore.Core

	coreConsole := NewCoreConsole(cfg.Console)
	if coreConsole != nil {
		cores = append(cores, coreConsole)
	}

	coreFile := NewCoreFile(cfg.File)
	if coreFile != nil {
		cores = append(cores, coreFile)
	}

	return zap.New(zapcore.NewTee(cores...))
}

func NewCoreConsole(cfg ConfigConsole) zapcore.Core {
	level, off := parseLevel(cfg.Level)
	if off {
		return nil
	}

	return zapcore.NewCore(
		zapcore.NewConsoleEncoder(log.NewEncoderConfig(!cfg.NoColor)),
		os.Stdout,
		level,
	)
}

func NewCoreFile(cfg ConfigFile) zapcore.Core {
	level, off := parseLevel(cfg.Level)
	if off {
		return nil
	}

	filename := cfg.Filename
	if filename != "" {
		filename, _ = filepath.Abs(filename)
	}
	if filename == "" {
		filename = filepath.Join(osx.ExeDir, osx.ExeName+".log")
	}

	writer := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    cfg.MaxSize,
		MaxAge:     cfg.MaxAge,
		MaxBackups: cfg.MaxBackups,
		Compress:   cfg.Compress,
		LocalTime:  true,
	}

	log.X.Named("log").Infof("save file %s", filename)

	return zapcore.NewCore(
		zapcore.NewConsoleEncoder(log.NewEncoderConfig(cfg.Color)),
		zapcore.AddSync(writer),
		level,
	)
}

func parseLevel(s string) (zapcore.Level, bool) {
	if s == "off" {
		return zapcore.InvalidLevel, true
	}

	level, err := zapcore.ParseLevel(s)
	if err != nil {
		level = zapcore.InfoLevel
	}

	return level, false
}
