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

func TestEnvLookup(t *testing.T) {
	cases := []struct {
		k  string
		v  SExpr
		ok bool
	}{
		{"foo", False, true},
		{"bar", True, true},
		{"baz", GetNil(), false},
	}

	p := NewEnv(nil)
	p.Define("bar", True)
	e := NewEnv(p)
	e.Define("foo", False)

	for _, tt := range cases {
		v, ok := e.Lookup(tt.k)

		if v != tt.v || ok != tt.ok {
			t.Errorf("%v.Lookup(%q) == (%s, %v), want (%s, %v)", e, tt.k, v, ok, tt.v, tt.ok)
		}
	}
}
