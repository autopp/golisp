package golisp

import (
	"errors"
	"fmt"
)

type Env struct {
	body map[string]SExpr
	prev *Env
}

func NewEnv(prev *Env) *Env {
	return &Env{make(map[string]SExpr), prev}
}

func (e *Env) Define(k string, v SExpr) error {
	if v, exists := e.body[k]; exists {
		return errors.New(fmt.Sprintf("%s is already defined with %s", k, v))
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
