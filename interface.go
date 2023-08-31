package goapp_gpython

import (
	"fmt"

	"golang.org/x/exp/slog"
)

var PyExecAtEnv func(func()) = func(f func()) {
	f()
}

var PyCallAtEnv func(func() any) any = func(f func() any) any {
	return f()
}

var PyLogWarn func(format string, a ...any) = func(format string, a ...any) {
	slog.Warn(fmt.Sprintf(format, a...))
}

var PyLogError func(format string, a ...any) = func(format string, a ...any) {
	slog.Error(fmt.Sprintf(format, a...))
}
