package env

import (
	"github.com/autopp/golisp/sexpr"
)

type Env struct {
	Prev *Env
	Body map[string]sexpr.SExpr
}
