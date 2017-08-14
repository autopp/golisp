package golisp

import (
	"testing"
)

func TestEqWithDifferentTypes(t *testing.T) {
	b := False
	n := Number(42)
	s := Symbol("foo")
	c := NewCons(Number(1), Number(2))
	cases := []struct {
		x, y SExpr
	}{
		{b, n},
		{b, s},
		{b, GetNil()},
		{b, c},
		{n, s},
		{n, GetNil()},
		{n, c},
		{s, GetNil()},
		{s, c},
		{GetNil(), c},
	}

	for _, tt := range cases {
		if Eq(tt.x, tt.y) {
			t.Errorf("Eq(%s, %s) == true, want false", tt.x, tt.y)
		}

		if Eq(tt.y, tt.x) {
			t.Errorf("Eq(%s, %s) == true, want false", tt.y, tt.x)
		}
	}
}
