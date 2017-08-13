package golisp

type Bool bool

const False = Bool(false)
const True = Bool(true)

func (sexpr Bool) IsNil() bool {
	return false
}

func (sexpr Bool) IsAtom() bool {
	return true
}

func (sexpr Bool) IsList() bool {
	return false
}

func (sexpr Bool) IsCons() bool {
	return false
}

func (sexpr Bool) Eq(other SExpr) bool {
	switch other.(type) {
	case Bool:
		return sexpr == other
	default:
		return false
	}
}
