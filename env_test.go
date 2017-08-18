package golisp

import (
	"testing"
)

func TestEnvDefine(t *testing.T) {
	cases := []struct {
		k   string
		v   SExpr
		err bool
	}{
		{"foo", False, false},
		{"bar", True, false},
		{"foo", True, true},
	}
	e := NewEnv(nil)

	for _, tt := range cases {
		err := e.Define(tt.k, tt.v)

		if tt.err {
			if err == nil {
				t.Errorf("%v.Define(%q, %s) == nil, want error", e, tt.k, tt.v)
			}
		} else {
			if err != nil {
				t.Errorf("%v.Define(%q, %s) == %s, want nil", e, tt.k, tt.v, err)
			}
		}
	}
}
