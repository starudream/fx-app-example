package app

import (
	"context"
	"errors"

	"go.uber.org/dig"

	"github.com/starudream/go-lib/v2/log"
	"github.com/starudream/go-lib/v2/x"
)

func Run(opts ...Option) func(cmd *Command, args []string) {
	return func(cmd *Command, args []string) {
		app := Default(cmd, opts...)
		handleError(app.Err())
		app.Run()
	}
}

func Action(fn func(cmd *Command, args []string) error, opts ...Option) func(cmd *Command, args []string) {
	return func(cmd *Command, args []string) {
		invoke := func(lc Lifecycle, sh Shutdowner, logger *log.SugaredLogger) {
			start := func(ctx context.Context) error {
				go func() {
					err := fn(cmd, args)
					if err != nil {
						logger.Error(err)
					}
					_ = sh.Shutdown(ExitCode(x.Ternary(err != nil, 1, 0)))
				}()
				return nil
			}
			lc.Append(StartHook(start))
		}
		app := Default(cmd, append(opts, Invoke(invoke))...)
		handleError(app.Err())
		app.Run()
	}
}

func handleError(err error) {
	if err == nil {
		return
	}

	logger := log.L().Named("app")

	re, pe := dig.RootCause(err), new(dig.PanicError)
	if errors.As(re, pe) {
		logger.Error(pe.Panic)
	}
	logger.Fatal(re)
}
