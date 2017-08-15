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

func (sexpr Bool) IsProc() bool {
	return false
}

func (sexpr Bool) String() string {
	if sexpr {
		return "#t"
	} else {
		return "#f"
	}
}
