package gpython_engine

import (
	"bytes"
	"goapp_commons/valconv"

	"github.com/go-python/gpython/py"
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

		buf := bytes.NewBuffer(nil)
		for i, v := range args {
			v, err := py.Str(v)
			if err != nil {
				return nil, err
			}

			rv, _ := P2G_Any(v, nil)
			buf.WriteString(valconv.AnyToString(rv))

			if i != len(args)-1 {
				buf.WriteString(string(sep))
			}
		}

		buf.WriteString(string(end))

		printf(buf.String())
		return py.None, nil
	}
}
