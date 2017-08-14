package golisp

import (
	"reflect"
)

type SExpr interface {
	IsNil() bool
	IsAtom() bool
	IsList() bool
	IsCons() bool
	Eq(SExpr) bool
	String() string
}

func Eq(x, y SExpr) bool {
	if reflect.TypeOf(x) != reflect.TypeOf(y) {
		return false
	}
	return true
}
