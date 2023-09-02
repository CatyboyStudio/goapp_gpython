package gpy_noc

import (
	"fmt"
	"goapp_gpython"
	"io/fs"
	"noc"

	"github.com/go-python/gpython/py"
	_ "github.com/go-python/gpython/stdlib"
	"github.com/pkg/errors"
)

type PyService struct {
	node       *noc.Node
	pyctx      py.Context
	mainModule *py.Module

	closeC chan bool
}

const KEY_PyService = "__pyservice"

func NodeGetService(n *noc.Node) *PyService {
	return noc.NodeGetService[*PyService](n, KEY_PyService)
}

func NewPyService(fs fs.FS) *PyService {
	opts := py.DefaultContextOpts()
	o := &PyService{
		pyctx:  py.NewContextWithFS(opts, fs),
		closeC: make(chan bool),
	}
	return o
}

func (ser *PyService) BindDefault(node *noc.Node) {
	node.BindService(KEY_PyService, ser)
}

func (ser *PyService) OnBindNode(node *noc.Node) {
	ser.node = node
}

func (ser *PyService) Dispose() {
	select {
	case <-ser.closeC:
		return
	default:
		close(ser.closeC)
		ser.pyctx.Close()
	}
}

func (ser *PyService) Context() py.Context {
	return ser.pyctx
}

func (ser *PyService) SetupModule(m *py.ModuleImpl) error {
	n := m.Info.Name
	mo, err := ser.pyctx.ModuleInit(m)
	if err != nil {
		return errors.Wrap(err, "SetupModule fail")
	}
	if n == "" {
		ser.mainModule = mo
	}
	return nil
}

func (ser *PyService) CompileAndRunScript(name string, script string) (any, error) {
	c, err := ser.CompileScript(name, script)
	if err != nil {
		return nil, err
	}
	return ser.RunCode(c)
}

func (ser *PyService) CompileScript(name string, script string) (*py.Code, error) {
	mode := py.ExecMode
	c, err := py.Compile(script, name, mode, 0, true)
	if err != nil {
		return nil, errors.Wrapf(err, "compile fail: %s", name)
	}
	return c, nil
}

func (ser *PyService) RunCode(code *py.Code) (any, error) {
	if code == nil {
		return nil, nil
	}
	if ser.mainModule == nil {
		pr := goapp_gpython.MakePrintFunc(func(msg string) {
			fmt.Println(msg)
		})
		mm := goapp_gpython.NewModule("", "")
		mm.Methods = append(mm.Methods, py.MustNewMethod("print", pr, 0, ""))
		err := ser.SetupModule(mm)
		if err != nil {
			return nil, err
		}
	}
	res, err := ser.pyctx.RunCode(code, ser.mainModule.Globals, ser.mainModule.Globals, nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return goapp_gpython.P2G_Value(res)
}
