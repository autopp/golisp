package golisp

type Nil struct{}

var nilV = &Nil{}

func GetNil() *Nil {
	return nilV
}

func (sexpr *Nil) IsNil() bool {
	return true
}

func (sexpr *Nil) IsAtom() bool {
	return true
}

func (sexpr *Nil) IsList() bool {
	return true
}

func (sexpr *Nil) IsCons() bool {
	return false
}

func (sexpr *Nil) Eq(other SExpr) bool {
	_, isNil := other.(*Nil)
	return isNil
}

func (sexpr *Nil) String() string {
	return "()"
}
