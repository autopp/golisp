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
	}

	for _, tt := range cases {
		e := NewEnv(map[string]SExpr{"foo": True}, nil)
		got, err := EvalSExpr(tt.in, e)

		if tt.err {
			if err == nil {
				t.Errorf("EvalSExpr(%s, %s) == (%s, nil), want error", tt.in, e, got)
			}
		} else {
			if got != tt.out || err != nil {
				t.Errorf("EvalSExpr(%s, %s) == (%s, %s), want (%s, nil)", tt.in, e, got, err, tt.out)
			}
		}
	}
}
