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

func NewCons(car, cdr SExpr) *Cons {
  return &Cons{car, cdr}
}

func (sexpr *Cons) IsNil() bool {
	return false
}

func (sexpr *Cons) IsAtom() bool {
	return false
}

func (sexpr *Cons) IsList() bool {
	return sexpr.Cdr.IsList()
}
