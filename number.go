package golisp

import "fmt"

type Number int

func (sexpr Number) IsNil() bool {
	return false
}

func (sexpr Number) IsAtom() bool {
	return true
}

func (sexpr Number) IsList() bool {
	return false
}

func (sexpr Number) IsCons() bool {
	return false
}

func (sexpr Number) String() string {
	return fmt.Sprintf("%d", sexpr)
}
