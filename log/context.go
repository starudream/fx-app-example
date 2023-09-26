package log

import (
	"context"
)

type Context struct {
	context.Context
}

type ctxkey struct{}

func WithContext(ctx context.Context, logger *SugaredLogger) *Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return &Context{context.WithValue(ctx, ctxkey{}, logger)}
}

func FromContext(ctx context.Context) *SugaredLogger {
	if ctx != nil {
		if logger, ok := ctx.Value(ctxkey{}).(*SugaredLogger); ok {
			return logger
		}
	}
	return L()
}
