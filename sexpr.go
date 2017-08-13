package golisp

type SExpr interface {
	IsNil() bool
	IsAtom() bool
	IsList() bool
	IsCons() bool
	Eq(SExpr) bool
}
