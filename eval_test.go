package golisp

import (
	"testing"
)

func TestEvalSExpr(t *testing.T) {
	cases := []struct {
		in, out SExpr
		err     bool
	}{
		{False, False, false},
		{True, True, false},
		{Number(42), Number(42), false},
		{GetNil(), GetNil(), false},
		{Symbol("foo"), True, false},
		{Symbol("bar"), GetNil(), true},
		{MakeList(Symbol("if"), True, Number(1)), Number(1), false},
		{MakeList(Symbol("if"), False, Number(1)), GetNil(), false},
		{MakeList(Symbol("if"), True, Number(1), Number(2)), Number(1), false},
		{MakeList(Symbol("if"), False, Number(1), Number(2)), Number(2), false},
		{MakeList(Symbol("quote"), Symbol("foo")), Symbol("foo"), false},
		{MakeList(Symbol("quote"), MakeList(Number(1), Number(2))), MakeList(Number(1), Number(2)), false},
	}

	for _, tt := range cases {
		e := NewGlobalEnv()
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
		}
	}
}
