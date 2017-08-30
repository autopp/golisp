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
	e := NewEnv(nil, nil)

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

	e := NewEnv(map[string]SExpr{"foo": False}, NewEnv(map[string]SExpr{"bar": True}, nil))

	for _, tt := range cases {
		v, ok := e.Lookup(tt.k)

		if v != tt.v || ok != tt.ok {
			t.Errorf("%v.Lookup(%q) == (%s, %v), want (%s, %v)", e, tt.k, v, ok, tt.v, tt.ok)
		}
	}
}

func TestEnvIsDefined(t *testing.T) {
	cases := []struct {
		in  string
		out bool
	}{}

	e := NewEnv(map[string]SExpr{"foo": False}, NewEnv(map[string]SExpr{"bar": True}, nil))

	for _, tt := range cases {
		got := e.IsDefined(tt.in)

		if got != tt.out {
			t.Errorf("%s.IsDefined(%q) == %v, want %v", e, tt.in, got, tt.out)
		}
	}
}
