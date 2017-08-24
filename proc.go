package golisp

import "fmt"

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

func (b *procBase) IsNil() bool {
	return false
}

func (b *procBase) IsAtom() bool {
	return false
}

func (b *procBase) IsList() bool {
	return false
}

func (b *procBase) IsCons() bool {
	return false
}

func (b *procBase) IsProc() bool {
	return true
}

type SpForm struct {
	*procBase
	Body func([]SExpr, *Env) (SExpr, error)
}

func NewSpForm(name string, required, optional int, body func([]SExpr, *Env) (SExpr, error)) *SpForm {
	return &SpForm{procBase: &procBase{name: name, required: required, optional: optional}, Body: body}
}

func (s *SpForm) String() string {
	return fmt.Sprintf("#<special %s>", s.Name())
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
