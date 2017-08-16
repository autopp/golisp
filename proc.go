package golisp

type Proc interface {
	Call(args []SExpr, env *Env) (SExpr, error)
	Name() string
	Required() int
	Optional() int
}

type procBase struct {
	name     string
	required int
	optional int
}

type SpForm struct {
	*procBase
}

type Func struct {
	*procBase
}

type BuiltinFunc struct {
	*Func
}

type UserFunc struct {
	*Func
}

func (b *procBase) Name() string {
	return b.name
}

func (b *procBase) Required() int {
	return b.required
}

func (b *procBase) Optional() int {
	return b.optional
}
