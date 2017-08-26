package golisp

import (
	"fmt"
)

type Env struct {
	body map[string]SExpr
	prev *Env
}

func NewEnv(body map[string]SExpr, prev *Env) *Env {
	if body == nil {
		body = make(map[string]SExpr)
	}
	return &Env{body, prev}
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
