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

func (b *procBase) Name() string {
	return b.name
}

func (b *procBase) Required() int {
	return b.required
}

func (b *procBase) Optional() int {
	return b.optional
}

type SpForm struct {
	*procBase
	Body func([]SExpr, *Env) (SExpr, error)
}

func NewSpForm(name string, required, optional int, body func([]SExpr, *Env) (SExpr, error)) *SpForm {
	return &SpForm{procBase: &procBase{name: name, required: required, optional: optional}, Body: body}
}

type Func struct {
	*procBase
}

type BuiltinFunc struct {
	*Func
	Body func([]SExpr, *Env) (SExpr, error)
}

type UserFunc struct {
	*Func
	ParamNames []string
	Body       SExpr
}
