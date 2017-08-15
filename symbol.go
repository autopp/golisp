package golisp

type Symbol string

func (sexpr Symbol) IsNil() bool {
	return false
}

func (sexpr Symbol) IsAtom() bool {
	return true
}

func (sexpr Symbol) IsList() bool {
	return false
}

func (sexpr Symbol) IsCons() bool {
	return false
}

func (sexpr Symbol) IsProc() bool {
	return false
}

func (sexpr Symbol) String() string {
	return string(sexpr)
}
