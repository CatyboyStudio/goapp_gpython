package gpython_engine

import "reflect"

type PE_Class struct {
}

func ForClassType(typ reflect.Type) *PE_Class {
	return nil
}

func ForClassName(fullname string) *PE_Class {
	return nil
}

func ForClassAny(value any) *PE_Class {
	return nil
}

func RegisterClass(cls *PE_Class) {

}

func OverwriteClass(cls *PE_Class) {

}

type PE_Func struct {
}

func ForFuncType(typ reflect.Type) *PE_Func {
	return nil
}

func ForFuncName(fullname string) *PE_Func {
	return nil
}

func RegisterFunc(f *PE_Func) {

}

func OverwriteFunc(cls *PE_Func) {

}

type PE_Factory struct {
}

type PE_Property struct {
}

type PE_Method struct {
}
