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
	case *Cons:
		if !v.IsList() {
			return GetNil(), fmt.Errorf("proc call should be list")
		}
		p, err := EvalSExpr(v.Car, e)
		if err != nil {
			return GetNil(), err
		}
		if !p.IsProc() {
			return GetNil(), fmt.Errorf("cannot call %s", p)
		}
		return p.(Proc).Call(v.Cdr.(*Cons).ToSlice(), e)
	default:
		return GetNil(), errors.New("not implemented type")
	}
}

func NewGlobalEnv() *Env {
	builtins := make(map[string]SExpr)
	builtins["if"] = NewSpForm("if", 2, 1, func(args []SExpr, env *Env) (SExpr, error) {
		c, err := EvalSExpr(args[0], env)
		if err != nil {
			return c, err
		}

		if c != False {
			return EvalSExpr(args[1], env)
		} else {
			if len(args) == 3 {
				return EvalSExpr(args[2], env)
			} else {
				return GetNil(), nil
			}
		}
	})

	builtins["quote"] = NewSpForm("quote", 1, 0, func(args []SExpr, env *Env) (SExpr, error) {
		return args[0], nil
	})
	return NewEnv(builtins, nil)
}
