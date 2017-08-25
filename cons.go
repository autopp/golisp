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

func (cons *Cons) IsNil() bool {
	return false
}

func (cons *Cons) IsAtom() bool {
	return false
}

func (cons *Cons) IsList() bool {
	return cons.Cdr.IsList()
}

func (cons *Cons) IsCons() bool {
	return true
}

func (cons *Cons) IsProc() bool {
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

func (cons *Cons) ToSlice() []SExpr {
	var l []SExpr
	for {
		l = append(l, cons.Car)
		if cons.Cdr.IsNil() {
			break
		}
		cons = cons.Cdr.(*Cons)
	}

	return l
}

func MakeList(elems ...SExpr) SExpr {
	var r SExpr = GetNil()

	for i := len(elems) - 1; i >= 0; i-- {
		r = NewCons(elems[i], r)
	}

	return r
}
