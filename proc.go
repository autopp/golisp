package golisp

type Proc interface {
	Call(args []SExpr, env *Env) (SExpr, error)
	Name() string
	Required() int
	Optional() int
}
