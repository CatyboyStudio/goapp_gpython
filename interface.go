package gpython_engine

import "log"

var PyExecAtEnv func(func()) = func(f func()) {
	f()
}

var PyCallAtEnv func(func() any) any = func(f func() any) any {
	return f()
}

var PyLogWarn func(format string, a ...any) = func(format string, a ...any) {
	log.Printf(format, a...)
	log.Println()
}

var PyLogError func(format string, a ...any) = func(format string, a ...any) {
	log.Printf(format, a...)
	log.Println()
}
