package golisp

type nilImpl struct{}

var nilV = &nilImpl{}

func GetNil() SExpr {
	return nilV
}

func (sexpr *nilImpl) IsNil() bool {
	return true
}

func (sexpr *nilImpl) IsAtom() bool {
	return true
}

func (sexpr *nilImpl) IsList() bool {
	return true
}

func (sexpr *nilImpl) IsCons() bool {
	return false
}

func (sexpr *nilImpl) Eq(other SExpr) bool {
	_, isNil := other.(*nilImpl)
	return isNil
}

func (sexpr *nilImpl) String() string {
	return "()"
}
