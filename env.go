package golisp

type Env struct {
	Prev *Env
	Body map[string]SExpr
}
