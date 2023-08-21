package gpython_engine

import "github.com/go-python/gpython/py"

type PyGoInt interface {
	GoInt() (int, error)
}
type PyGoInt64 interface {
	GoInt64() (int64, error)
}

func P2G_Any(p py.Object, unknow func(py.Object) any) (any, error) {
	if p == nil {
		return nil, nil
	}
	switch rv := p.(type) {
	case py.NoneType:
		return nil, nil
	case py.Bool:
		return (bool)(rv), nil
	case py.Float:
		return (float64)(rv), nil
	case py.String:
		return (string)(rv), nil
	case py.StringDict:
		r := map[string]any{}
		for k, v := range rv {
			v2, err := P2G_Any(v, unknow)
			if err != nil {
				return nil, err
			}
			r[k] = v2
		}
		return r, nil
	case py.Tuple:
		var r []any
		for _, v := range rv {
			v2, err := P2G_Any(v, unknow)
			if err != nil {
				return nil, err
			}
			r = append(r, v2)
		}
		return r, nil
	case *py.List:
		var r []any
		for _, v := range rv.Items {
			v2, err := P2G_Any(v, unknow)
			if err != nil {
				return nil, err
			}
			r = append(r, v2)
		}
		return r, nil
	}
	if f, ok := p.(PyGoInt); ok {
		return f.GoInt()
	}
	if f, ok := p.(PyGoInt64); ok {
		return f.GoInt64()
	}
	if unknow == nil {
		return nil, nil
	}
	return unknow(p), nil
}
