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

	if x == y {
		return true
	}

	switch z := x.(type) {
	case *Nil:
		return true
	case *Cons:
		w := y.(*Cons)
		return Eq(z.Car, w.Car) && Eq(z.Cdr, w.Cdr)
	}

	return false
}
