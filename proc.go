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

func (s *SpForm) Call(args []SExpr, env *Env) (SExpr, error) {
	return s.Body(args, env)
}

type Func struct {
	*procBase
}

func (f *Func) String() string {
	if f.Name() == "" {
		return "#<lambda>"
	}
	return fmt.Sprintf("#<lambda %s>", f.Name())
}

type BuiltinFunc struct {
	*Func
	Body func([]SExpr, *Env) (SExpr, error)
}

func (f *Func) evalArgs(args []SExpr, env *Env) ([]SExpr, error) {
	r := make([]SExpr, len(args))

	for i, a := range args {
		x, err := EvalSExpr(a, env)
		if err != nil {
			return nil, err
		}
		r[i] = x
	}
	return r, nil
}

func NewBuiltinFunc(name string, required, optional int, body func([]SExpr, *Env) (SExpr, error)) *BuiltinFunc {
	return &BuiltinFunc{Func: &Func{procBase: &procBase{name: name, required: required, optional: optional}}, Body: body}
}

func (f *BuiltinFunc) Call(args []SExpr, env *Env) (SExpr, error) {
	args, err := f.evalArgs(args, env)

	if err != nil {
		return GetNil(), err
	}

	return f.Body(args, env)
}

type UserFunc struct {
	*Func
	ParamNames []string
	Body       SExpr
}

func NewUserFunc(name string, paramNames []string, body SExpr) *UserFunc {
	return &UserFunc{Func: &Func{procBase: &procBase{name: name, required: len(paramNames), optional: 0}}, ParamNames: paramNames, Body: body}
}

func (f *UserFunc) Call(args []SExpr, env *Env) (SExpr, error) {
	args, err := f.evalArgs(args, env)

	if err != nil {
		return GetNil(), err
	}

	m := make(map[string]SExpr, len(args))
	for i := 0; i < len(args); i++ {
		m[f.ParamNames[i]] = args[i]
	}

	e := NewEnv(m, env)
	return EvalSExpr(f.Body, e)
}
