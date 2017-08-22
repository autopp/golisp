package golisp

import (
	"errors"
	"fmt"
)

func EvalSExpr(s SExpr, e *Env) (SExpr, error) {
	switch v := s.(type) {
	case Bool, Number, *Nil:
		return v, nil
	case Symbol:
		val, ok := e.Lookup(string(v))
		if ok {
			return val, nil
		}
		return GetNil(), fmt.Errorf("%s is not defined", v)
	default:
		return GetNil(), errors.New("not implemented type")
	}
}
