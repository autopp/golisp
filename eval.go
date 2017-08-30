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

		proc := p.(Proc)

		var args []SExpr
		var argc int
		if v.Cdr.IsNil() {
			argc = 0
			args = make([]SExpr, 0)
		} else {
			args = v.Cdr.(*Cons).ToSlice()
			argc = len(args)
		}

		if proc.Optional() < 0 {
			if argc < proc.Required() {
				return GetNil(), fmt.Errorf("got %d arguments, want %d or more", argc, proc.Required())
			}
		} else {
			if argc < proc.Required() || argc > proc.Required()+proc.Optional() {
				if proc.Optional() == 0 {
					return GetNil(), fmt.Errorf("got %d arguments, want %d", argc, proc.Required())
				}
				return GetNil(), fmt.Errorf("got %d arguments, want between %d to %d", argc, proc.Required(), proc.Required()+proc.Optional())
			}
		}

		return proc.Call(args, e)
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
			}
			return GetNil(), nil
		}
	})

	builtins["quote"] = NewSpForm("quote", 1, 0, func(args []SExpr, env *Env) (SExpr, error) {
		return args[0], nil
	})

	builtins["define"] = NewSpForm("define", 2, 0, func(args []SExpr, env *Env) (SExpr, error) {
		s, ok := args[0].(Symbol)

		if !ok {
			return GetNil(), errors.New("define: want symbol at 1st argument")
		}

		if env.IsDefined(string(s)) {
			return GetNil(), fmt.Errorf("%s is already defined at current scope", s)
		}

		env.Define(string(s), args[1])

		return s, nil
	})

	builtins["cons"] = NewBuiltinFunc("cons", 2, 0, func(args []SExpr, env *Env) (SExpr, error) {
		return NewCons(args[0], args[1]), nil
	})

	builtins["car"] = NewBuiltinFunc("car", 1, 0, func(args []SExpr, env *Env) (SExpr, error) {
		x, ok := args[0].(*Cons)

		if !ok {
			return nil, fmt.Errorf("car: got %s, want cons", args[0])
		}

		return x.Car, nil
	})

	builtins["cdr"] = NewBuiltinFunc("cdr", 1, 0, func(args []SExpr, env *Env) (SExpr, error) {
		x, ok := args[0].(*Cons)

		if !ok {
			return nil, fmt.Errorf("car: got %s, want cons", args[0])
		}

		return x.Cdr, nil
	})

	builtins["null"] = NewBuiltinFunc("null", 1, 0, func(args []SExpr, env *Env) (SExpr, error) {
		return Bool(args[0].IsNil()), nil
	})

	builtins["eq?"] = NewBuiltinFunc("eq?", 2, 0, func(args []SExpr, env *Env) (SExpr, error) {
		return Bool(args[0] == args[1]), nil
	})

	builtins["equal?"] = NewBuiltinFunc("eq?", 2, 0, func(args []SExpr, env *Env) (SExpr, error) {
		return Bool(Eq(args[0], args[1])), nil
	})

	builtins["+"] = NewBuiltinFunc("+", 0, -1, func(args []SExpr, env *Env) (SExpr, error) {
		r := 0
		for i, x := range args {
			n, ok := x.(Number)
			if !ok {
				return GetNil(), fmt.Errorf("%d argument is not a number", i+1)
			}
			r += int(n)
		}

		return Number(r), nil
	})
	return NewEnv(builtins, nil)
}
