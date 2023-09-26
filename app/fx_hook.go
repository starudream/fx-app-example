package app

import (
	"go.uber.org/fx"
)

type (
	Hook = fx.Hook

	HookFunc = fx.HookFunc
)

func StartHook[T HookFunc](start T) Hook {
	return fx.StartHook(start)
}

func StopHook[T HookFunc](stop T) Hook {
	return fx.StopHook(stop)
}

func StartStopHook[T1 HookFunc, T2 HookFunc](start T1, stop T2) Hook {
	return fx.StartStopHook(start, stop)
}
