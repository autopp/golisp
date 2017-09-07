package golisp

import (
	"fmt"
	"io"
	"os"
)

type Env struct {
	body   map[string]SExpr
	prev   *Env
	output io.Writer
}

func NewEnv(body map[string]SExpr, prev *Env) *Env {
	return NewEnvWithOutput(body, prev, nil)
}

func NewEnvWithOutput(body map[string]SExpr, prev *Env, output io.Writer) *Env {
	if body == nil {
		body = make(map[string]SExpr)
	}

	if output == nil {
		if prev != nil {
			output = prev.Output()
		} else {
			output = os.Stdout
		}
	}

	return &Env{body, prev, output}
}

func (e *Env) Output() io.Writer {
	return e.output
}

func (e *Env) Define(k string, v SExpr) error {
	if w, exists := e.body[k]; exists {
		return fmt.Errorf("%s is already defined with %s", k, w)
	}
	e.body[k] = v
	return nil
}

func (e *Env) Lookup(k string) (SExpr, bool) {
	v, ok := e.body[k]

	if ok {
		return v, true
	} else {
		if e.prev != nil {
			return e.prev.Lookup(k)
		} else {
			return GetNil(), false
		}
	}
}

func (e *Env) IsDefined(k string) bool {
	_, ok := e.body[k]
	return ok
}
