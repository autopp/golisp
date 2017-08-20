package golisp

import (
	"errors"
)

func EvalSExpr(s SExpr, e *Env) (SExpr, error) {
	switch v := s.(type) {
	case Bool, Number, *Nil:
		return v, nil
	default:
		return GetNil(), errors.New("not implemented type")
	}
}
