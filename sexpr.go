package golisp

type SExpr interface {
}

type Bool bool
type Number int
type Symbol string
type Nil struct {}
type Cons struct {
  Car, Cdr SExpr
}
type Proc struct {}
