package golisp

import (
	"bytes"
)

type Cons struct {
	Car, Cdr SExpr
}

func NewCons(car, cdr SExpr) *Cons {
	return &Cons{car, cdr}
}

func (sexpr *Cons) IsNil() bool {
	return false
}

func (sexpr *Cons) IsAtom() bool {
	return false
}

func (sexpr *Cons) IsList() bool {
	return sexpr.Cdr.IsList()
}

func (sexpr *Cons) IsCons() bool {
	return true
}

func (sexpr *Cons) IsProc() bool {
	return false
}

func (cons *Cons) String() string {
	buf := bytes.NewBufferString("(")
	buf.WriteString(cons.Car.String())
	cdr := cons.Cdr

	for end := false; !end; {
		switch v := cdr.(type) {
		case *Cons:
			buf.WriteString(" ")
			buf.WriteString(v.Car.String())
			cdr = v.Cdr
		case *Nil:
			end = true
		default:
			buf.WriteString(" . ")
			buf.WriteString(v.String())
			end = true
		}
	}
	buf.WriteString(")")
	return buf.String()
}

func MakeList(elems ...SExpr) SExpr {
	var r SExpr = GetNil()

	for i := len(elems) - 1; i >= 0; i-- {
		r = NewCons(elems[i], r)
	}

	return r
}
