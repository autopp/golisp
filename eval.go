package golisp

import (
	"errors"
	"fmt"
	"io"
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
		args := ToSlice(v.Cdr)
		return applyProc(proc, args, e)
	default:
		return GetNil(), errors.New("not implemented type")
	}
}

func NewGlobalEnv() *Env {
	return NewGlobalEnvWithOutput(nil)
}

func NewGlobalEnvWithOutput(output io.Writer) *Env {
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

	builtins["lambda"] = NewSpForm("lambda", 2, 0, func(args []SExpr, env *Env) (SExpr, error) {
		if !args[0].IsList() {
			return GetNil(), errors.New("lambda: 1st argument shoud be a list of parameter")
		}

		s := ToSlice(args[0])
		p := make([]string, len(s))

		for i := 0; i < len(s); i++ {
			if n, ok := s[i].(Symbol); !ok {
				return GetNil(), errors.New("lambda: 1st argument shoud be a list of parameter")
			} else {
				p[i] = string(n)
			}
		}

		return NewUserFunc("", p, args[1], env), nil
	})

	builtins["define"] = NewSpForm("define", 2, 0, func(args []SExpr, env *Env) (SExpr, error) {
		var v SExpr
		var err error

		s, ok := args[0].(Symbol)
		if ok {
			v, err = EvalSExpr(args[1], env)
			if err != nil {
				return GetNil(), err
			}
		} else {
			if !args[0].IsList() || args[0].IsNil() {
				return GetNil(), fmt.Errorf("define: 1 st argument shoud be a symbol or list of symbol")
			}

			sl := ToSlice(args[0])
			p := make([]string, len(sl))

			for i := 0; i < len(sl); i++ {
				if n, ok := sl[i].(Symbol); !ok {
					return GetNil(), errors.New("define: 1 st argument shoud be a symbol or list of symbol")
				} else {
					p[i] = string(n)
				}
			}

			s = sl[0].(Symbol)
			v = NewUserFunc(p[0], p[1:], args[1], env)
		}

		if env.IsDefined(string(s)) {
			return GetNil(), fmt.Errorf("define: %s is already defined at current scope", s)
		}
		env.Define(string(s), v)

		return s, nil
	})

	builtins["begin"] = NewSpForm("begin", 1, -1, func(args []SExpr, env *Env) (SExpr, error) {
		for i := 0; i < len(args)-1; i++ {
			_, err := EvalSExpr(args[i], env)

			if err != nil {
				return GetNil(), err
			}
		}

		return EvalSExpr(args[len(args)-1], env)
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

	builtins["eval"] = NewBuiltinFunc("eval", 1, 0, func(args []SExpr, env *Env) (SExpr, error) {
		return EvalSExpr(args[0], env)
	})

	builtins["apply"] = NewBuiltinFunc("apply", 1, -1, func(args []SExpr, env *Env) (SExpr, error) {
		f, ok := args[0].(Proc)

		if !ok || !f.IsFunc() {
			return GetNil(), errors.New("apply: 1st argument shoud be function")
		}

		r, err := applyProc(f, args[1:], env)

		if err != nil {
			return r, errors.New("apply: error occured: " + err.Error())
		}
		return r, nil
	})

	builtins["display"] = NewBuiltinFunc("display", 1, 0, func(args []SExpr, env *Env) (SExpr, error) {
		fmt.Fprintln(env.Output(), args[0])
		return GetNil(), nil
	})

	return NewEnvWithOutput(builtins, nil, output)
}

func applyProc(proc Proc, args []SExpr, env *Env) (SExpr, error) {
	argc := len(args)

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

	return proc.Call(args, env)
}
