package goapp_gpython

import "github.com/go-python/gpython/py"

func NewModule(name string, script string) *py.ModuleImpl {
	m := &py.ModuleImpl{
		CodeSrc: script,
		Globals: py.NewStringDict(),
	}
	m.Info.Name = name
	return m
}
