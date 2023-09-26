package x

import (
	"reflect"
)

func Must0(err any, messageArgs ...any) {
	must(err, messageArgs...)
}

func Must1[T any](val T, err any, messageArgs ...any) T {
	must(err, messageArgs...)
	return val
}

func Must2[T1 any, T2 any](val1 T1, val2 T2, err any, messageArgs ...any) (T1, T2) {
	must(err, messageArgs...)
	return val1, val2
}

func Must3[T1 any, T2 any, T3 any](val1 T1, val2 T2, val3 T3, err any, messageArgs ...any) (T1, T2, T3) {
	must(err, messageArgs...)
	return val1, val2, val3
}

func must(err any, args ...any) {
	if err == nil {
		return
	}
	switch e := err.(type) {
	case bool:
		if !e {
			message := MessageFromMsgAndArgs(args...)
			if message == "" {
				message = "not ok"
			}
			panic(message)
		}
	case error:
		message := MessageFromMsgAndArgs(args...)
		if message != "" {
			panic(message + ": " + e.Error())
		} else {
			panic(e.Error())
		}
	default:
		panic("must: invalid err type '" + reflect.TypeOf(err).Name() + "', should either be a bool or an error")
	}
}
