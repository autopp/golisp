package golisp

type SExpr interface {
	IsNil() bool
	IsAtom() bool
	IsLIst() bool
}

type Bool bool
type Number int
type Symbol string
type Nil struct{}
type Cons struct {
	Car, Cdr SExpr
}

func (sexpr Bool) IsNil() bool {
	return false
}

func (sexpr Bool) IsAtom() bool {
	return true
}

func (sexpr Bool) IsLIst() bool {
	return false
}

func (sexpr Number) IsNil() bool {
	return false
}

func (sexpr Number) IsAtom() bool {
	return true
}

func (sexpr Number) IsLIst() bool {
	return false
}

func (sexpr Symbol) IsNil() bool {
	return false
}

func (sexpr Symbol) IsAtom() bool {
	return true
}

func (sexpr Symbol) IsLIst() bool {
	return false
}

func (sexpr *Nil) IsNil() bool {
	return false
}

func (sexpr *Nil) IsAtom() bool {
	return true
}

func (sexpr *Nil) IsList() bool {
	return true
}
