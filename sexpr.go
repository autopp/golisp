package golisp

type SExpr interface {
  isNil() bool
  isAtom() bool
  isList() bool
}

type Bool bool
type Number int
type Symbol string
type Nil struct {}
type Cons struct {
  Car, Cdr SExpr
}

func (sexpr Bool)isNil() bool {
  return false
}

func (sexpr Bool)isAtom() bool {
  return true
}

func (sexpr Bool)isList() bool {
  return false
}

func (sexpr Number)isNil() bool {
  return false
}

func (sexpr Number)isAtom() bool {
  return true
}

func (sexpr Number)isList() bool {
  return false
}

func (sexpr Symbol)isNil() bool {
  return false
}

func (sexpr Symbol)isAtom() bool {
  return true
}

func (sexpr Symbol)isList() bool {
  return false
}

func (sexpr *Nil)isNil() bool {
  return false
}

func (sexpr *Nil)isAtom() bool {
  return true
}

func (sexpr *Nil)isList() bool {
  return true
}
