package gpython_engine

import (
	"github.com/go-python/gpython/py"
	"github.com/gookit/goutil/byteutil"
)

func MakePrintFunc(printf func(msg string)) func(py.Object, py.Tuple, py.StringDict) (py.Object, error) {
	return func(self py.Object, args py.Tuple, kwargs py.StringDict) (py.Object, error) {
		var (
			sepObj py.Object = py.String(" ")
			endObj py.Object = py.String("")
			err    error
		)
		kwlist := []string{"sep", "end"}
		err = py.ParseTupleAndKeywords(nil, kwargs, "|ss:print", kwlist, &sepObj, &endObj)
		if err != nil {
			return nil, err
		}
		sep := sepObj.(py.String)
		end := endObj.(py.String)

		buf := byteutil.NewBuffer()
		for i, v := range args {
			v, err := py.Str(v)
			if err != nil {
				return nil, err
			}

			rv, _ := P2G_Any(v, nil)
			buf.WriteAny(rv)

			if i != len(args)-1 {
				buf.WriteAny(sep)
			}
		}

		buf.WriteAny(end)

		printf(buf.String())
		return py.None, nil
	}
}
