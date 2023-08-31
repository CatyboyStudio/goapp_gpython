package goapp_gpython

import (
	"goapp_commons/valconv"
	"math/big"

	"github.com/go-python/gpython/py"
)

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

func P2G_Value(p py.Object) (any, error) {
	return P2G_Any(p, func(v py.Object) any {
		return valconv.AnyToString(v)
	})
}

func G2P_Value(v any) py.Object {
	if v == nil {
		return py.None
	}
	switch rv := v.(type) {
	case bool:
		if rv {
			return py.True
		} else {
			return py.False
		}
	case int:
		return py.Int(rv)
	case uint8:
		return py.Int(int(rv))
	case int16:
		return py.Int(int(rv))
	case uint16:
		return py.Int(int(rv))
	case int32:
		return py.Int(int(rv))
	case uint32:
		return py.Int(int(rv))
	case int64:
		o := &big.Int{}
		o.SetInt64(rv)
		return (*py.BigInt)(o)
	case uint64:
		o := &big.Int{}
		o.SetUint64(rv)
		return (*py.BigInt)(o)
	case float32:
		return py.Float(float64(rv))
	case float64:
		return py.Float(rv)
	case string:
		return py.String(rv)
	case []any:
		tuple := make([]py.Object, len(rv))
		for i, v := range rv {
			tuple[i] = G2P_Value(v)
		}
		return py.NewListFromItems(tuple)
	case map[string]any:
		dict := py.NewStringDict()
		for k, v := range rv {
			dict[k] = G2P_Value(v)
		}
		return dict
	}
	return py.None
}
