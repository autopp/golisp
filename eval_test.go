package golisp

import (
	"bytes"
	"testing"
)

func quote(s SExpr) SExpr {
	return MakeList(Symbol("quote"), s)
}

func begin(s ...SExpr) SExpr {
	return MakeList(append([]SExpr{Symbol("begin")}, s...)...)
}

var boolCons = quote(NewCons(False, True))

func TestEvalSExpr(t *testing.T) {
	cases := []struct {
		in, out SExpr
		printed string
		err     bool
	}{
		{False, False, "", false},
		{True, True, "", false},
		{Number(42), Number(42), "", false},
		{GetNil(), GetNil(), "", false},
		{Symbol("foo"), True, "", false},
		{Symbol("bar"), GetNil(), "", true},
		{MakeList(Number(1), Number(2), Number(3)), GetNil(), "", true},
		{MakeList(Symbol("bar"), Number(1), Number(2)), GetNil(), "", true},
		{MakeList(Symbol("if"), True, Number(1)), Number(1), "", false},
		{MakeList(Symbol("if"), False, Number(1)), GetNil(), "", false},
		{MakeList(Symbol("if"), True, Number(1), Number(2)), Number(1), "", false},
		{MakeList(Symbol("if"), False, Number(1), Number(2)), Number(2), "", false},
		{NewCons(Symbol("if"), True), GetNil(), "", true},
		{MakeList(Symbol("if"), True), GetNil(), "", true},
		{MakeList(Symbol("if"), True, Number(1), Number(2), Number(3)), GetNil(), "", true},
		{MakeList(Symbol("quote"), Symbol("foo")), Symbol("foo"), "", false},
		{MakeList(Symbol("quote"), MakeList(Number(1), Number(2))), MakeList(Number(1), Number(2)), "", false},
		{MakeList(MakeList(Symbol("lambda"), MakeList(Symbol("x")), MakeList(Symbol("+"), Number(1), Symbol("x"))), Number(41)), Number(42), "", false},
		{MakeList(MakeList(MakeList(Symbol("lambda"), MakeList(Symbol("x")), MakeList(Symbol("lambda"), MakeList(Symbol("y")), MakeList(Symbol("+"), Symbol("x"), Symbol("y")))), Number(1)), Number(41)), Number(42), "", false},
		{MakeList(Symbol("lambda"), NewCons(Symbol("x"), Symbol("y")), MakeList(Symbol("+"), Number(1), Symbol("x"))), GetNil(), "", true},
		{MakeList(Symbol("lambda"), MakeList(Symbol("x"), Number(1)), MakeList(Symbol("+"), Number(1), Symbol("x"))), GetNil(), "", true},
		{MakeList(Symbol("define"), Symbol("x"), Number(42)), Symbol("x"), "", false},
		{MakeList(Symbol("define"), MakeList(Symbol("f"), Symbol("x")), MakeList(Symbol("+"), Symbol("x"), Number(1))), Symbol("f"), "", false},
		{MakeList(Symbol("define"), Symbol("foo"), Number(42)), GetNil(), "", true},
		{MakeList(Symbol("begin"), MakeList(Symbol("define"), Symbol("x"), Number(41)), MakeList(Symbol("+"), Symbol("x"), Number(1))), Number(42), "", false},
		{
			MakeList(
				Symbol("begin"),
				MakeList(Symbol("define"), MakeList(Symbol("f"), Symbol("x")), MakeList(Symbol("+"), Symbol("x"), Number(1))),
				MakeList(Symbol("f"), Number(41))),
			Number(42), "", false,
		},
		{MakeList(Symbol("begin"), Symbol("x"), MakeList(Symbol("+"), Symbol("x"), Number(1))), GetNil(), "", true},
		{MakeList(Symbol("cons"), False, True), NewCons(False, True), "", false},
		{MakeList(Symbol("car"), quote(NewCons(False, True))), False, "", false},
		{MakeList(Symbol("car"), GetNil()), GetNil(), "", true},
		{MakeList(Symbol("cdr"), quote(NewCons(False, True))), True, "", false},
		{MakeList(Symbol("cdr"), GetNil()), GetNil(), "", true},
		{MakeList(Symbol("null"), GetNil()), True, "", false},
		{MakeList(Symbol("null"), Number(42)), False, "", false},
		{MakeList(Symbol("eq?"), Number(1), Number(1)), True, "", false},
		{MakeList(Symbol("eq?"), Number(1), Number(2)), False, "", false},
		{MakeList(Symbol("eq?"), GetNil(), GetNil()), True, "", false},
		{MakeList(Symbol("eq?"), boolCons, boolCons), True, "", false},
		{MakeList(Symbol("eq?"), boolCons, quote(NewCons(False, True))), False, "", false},
		{MakeList(Symbol("equal?"), Number(1), Number(1)), True, "", false},
		{MakeList(Symbol("equal?"), Number(1), Number(2)), False, "", false},
		{MakeList(Symbol("equal?"), GetNil(), GetNil()), True, "", false},
		{MakeList(Symbol("equal?"), boolCons, boolCons), True, "", false},
		{MakeList(Symbol("equal?"), boolCons, quote(NewCons(False, True))), True, "", false},
		{MakeList(Symbol("+")), Number(0), "", false},
		{MakeList(Symbol("+"), Number(1)), Number(1), "", false},
		{MakeList(Symbol("+"), Number(41), Number(1)), Number(42), "", false},
		{MakeList(Symbol("+"), Number(41), quote(Symbol("a"))), GetNil(), "", true},
		{
			begin(
				MakeList(Symbol("define"), Symbol("x"), quote(MakeList(Symbol("+"), Number(1), Number(41)))),
				MakeList(Symbol("eval"), Symbol("x"))),
			Number(42), "", false,
		},
		{MakeList(Symbol("apply"), Symbol("cons"), Number(1), Number(2)), NewCons(Number(1), Number(2)), "", false},
		{MakeList(Symbol("apply"), Symbol("cons"), Number(1)), GetNil(), "", true},
		{MakeList(Symbol("apply"), Symbol("if"), True, Number(42)), GetNil(), "", true},
		{MakeList(Symbol("display"), quote(MakeList(Number(1), Number(2), Number(3)))), GetNil(), "(1 2 3)\n", false},
	}

	for _, tt := range cases {
		buf := new(bytes.Buffer)
		e := NewGlobalEnvWithOutput(buf)
		e.Define("foo", True)
		got, err := EvalSExpr(tt.in, e)

		if tt.err {
			if err == nil {
				t.Errorf("EvalSExpr(%s, %s) == (%s, nil), want error", tt.in, e, got)
			}
		} else {
			if !Eq(got, tt.out) || err != nil {
				t.Errorf("EvalSExpr(%v, %v) == (%v, %v), want (%s, nil)", tt.in, e, got, err, tt.out)
			}
			if buf.String() != tt.printed {
				t.Errorf("EvalSExpr(%v, %v) output %q, want %q", tt.in, e, buf.String(), tt.printed)
			}
		}
	}
}
