package golisp

import (
	"reflect"
)

type SExpr interface {
	IsNil() bool
	IsAtom() bool
	IsList() bool
	IsCons() bool
	IsProc() bool
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

func ToSlice(x SExpr) []SExpr {
	var s []SExpr
	for ; !x.IsNil(); x = x.(*Cons).Cdr {
		s = append(s, x.(*Cons).Car)
	}

	return s
}
