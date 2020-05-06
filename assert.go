package struct_copy

import "reflect"

// 基础断言
type AssertInterface interface {
	IsNil(obj interface{}, message string)
	NotNil(obj interface{}, message string)
	NotEmpty(obj interface{}, message string)
	IsTrue(expression bool, message string)
}

type Assert struct {
}

func NewAssert() *Assert {
	return &Assert{}
}

func (at *Assert) IsNil(obj interface{}, message string) {
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Ptr {
		if !v.IsNil() {
			panic(message)
		}
	} else {
		if obj != nil {
			panic(message)
		}
	}
}

func (at *Assert) NotNil(obj interface{}, message string) {
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			panic(message)
		}
	} else {
		if obj == nil {
			panic(message)
		}
	}
}

func (at *Assert) NotEmpty(obj interface{}, message string) {
	// todo
}

func (at *Assert) IsTrue(expression bool, message string) {
	if !expression {
		panic(message)
	}
}
