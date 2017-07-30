package sexpr

type SExpr interface {
	IsNil() bool
	IsAtom() bool
	IsList() bool
}

type Bool bool
type Number int
type Symbol string
type nilImpl struct{}
type Cons struct {
	Car, Cdr SExpr
}

func (sexpr Bool) IsNil() bool {
	return false
}

func (sexpr Bool) IsAtom() bool {
	return true
}

func (sexpr Bool) IsList() bool {
	return false
}

func (sexpr Number) IsNil() bool {
	return false
}

func (sexpr Number) IsAtom() bool {
	return true
}

func (sexpr Number) IsList() bool {
	return false
}

func (sexpr Symbol) IsNil() bool {
	return false
}

func (sexpr Symbol) IsAtom() bool {
	return true
}

func (sexpr Symbol) IsList() bool {
	return false
}

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