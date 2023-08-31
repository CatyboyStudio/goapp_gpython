package goapp_gpython

import (
	"errors"

	"github.com/go-python/gpython/py"
)

// Class: bool, Kind: Value
func P2G_Bool(v py.Object) bool {
	val, err := py.MakeBool(v)
	if err != nil {
		PyLogWarn("P2G_Bool fail %v: %v\n", v, err)
		return false
	}
	return bool(val.(py.Bool))
}

func G2P_Bool(v bool) py.Object {
	return py.NewBool(v)
}

// Class: int, Kind: Value
func P2G_Int(v py.Object) int {
	val, err := py.MakeGoInt(v)
	if err != nil {
		PyLogWarn("P2G_Int fail %v: %v\n", v, err)
		return 0
	}
	return val
}
func G2P_Int(v int) py.Object {
	return py.Int(v)
}

// Class: int64, Kind: Value
func P2G_Int64(v py.Object) int64 {
	val, err := py.MakeGoInt64(v)
	if err != nil {
		PyLogWarn("P2G_Int64 fail %v: %v\n", v, err)
		return 0
	}
	return val
}
func G2P_Int64(v int64) py.Object {
	return py.Int(v)
}

// Class: float64, Kind: Value
func P2G_Float64(v py.Object) float64 {
	val, err := py.MakeFloat(v)
	if err != nil {
		PyLogWarn("P2G_Float fail %v: %v\n", v, err)
		return 0
	}
	return float64(val.(py.Float))
}
func G2P_Float64(v float64) py.Object {
	return py.Float(v)
}

// Class: string, Kind: Value
func P2G_String(v py.Object) string {
	if v, err := py.StrAsString(v); err == nil {
		return v
	}
	return ""
}

func G2P_String(v string) py.Object {
	return py.String(v)
}

// Class: error
func P2G_Error(v py.Object) error {
	if o, ok := v.(*py.Exception); ok {
		return o
	}
	val := P2G_String(v)
	return errors.New(val)
}

func G2P_Error(v error) py.Object {
	if v == nil {
		return py.None
	}
	return py.ExceptionNewf(py.RuntimeError, "%v", v)
}
