package gpy_noc

import (
	"goapp_gpython"

	"github.com/go-python/gpython/py"
)

type ServiceBuilder struct {
	modules map[string]*py.ModuleImpl
}

func NewServiceBuilder() *ServiceBuilder {
	return &ServiceBuilder{
		modules: make(map[string]*py.ModuleImpl),
	}
}

func (b *ServiceBuilder) GetModule(n string) *py.ModuleImpl {
	return b.modules[n]
}

func (b *ServiceBuilder) MustModule(n string) *py.ModuleImpl {
	if m, ok := b.modules[n]; ok {
		return m
	}
	m := goapp_gpython.NewModule(n, "")
	b.modules[n] = m
	return m
}

func (b *ServiceBuilder) Build(s *PyService) error {
	for _, m := range b.modules {
		err := s.SetupModule(m)
		if err != nil {
			return err
		}
	}
	return nil
}
