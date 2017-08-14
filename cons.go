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

func (sexpr *Cons) Eq(other SExpr) bool {
	cons, ok := other.(*Cons)
	if !ok {
		return false
	}

	return sexpr.Car.Eq(cons.Car) && sexpr.Cdr.Eq(cons.Cdr)
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
