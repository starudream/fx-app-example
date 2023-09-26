package app

import (
	"time"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"

	"github.com/starudream/go-lib/v2/config"
	"github.com/starudream/go-lib/v2/log"
	"github.com/starudream/go-lib/v2/x"
)

type (
	Option = fx.Option

	Lifecycle  = fx.Lifecycle
	Shutdowner = fx.Shutdowner
)

var (
	Module  = fx.Module
	Options = fx.Options

	Supply  = fx.Supply
	Provide = fx.Provide
	Invoke  = fx.Invoke

	ExitCode = fx.ExitCode
)

func Default(cmd *Command, opts ...Option) *fx.App {
	opts = append(opts,
		fx.StartTimeout(time.Minute),
		fx.StopTimeout(time.Minute),
		fx.RecoverFromPanics(),
		Supply(cmd),
		Provide(config.LoadFromCommand),
		Provide(log.NewFromKoanf),
		fx.WithLogger(wrapLogger),
	)
	return fx.New(opts...)
}

func wrapLogger(k *config.Koanf, logger *log.Logger) fxevent.Logger {
	return x.Ternary[fxevent.Logger](config.Debug(k), &fxevent.ZapLogger{Logger: logger.Named("fx")}, fxevent.NopLogger)
}
