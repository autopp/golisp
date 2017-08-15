package golisp

import (
	"testing"
)

func TestBoolIsNil(t *testing.T) {
	cases := []struct {
		in  SExpr
		out bool
	}{
		{False, false},
		{True, false},
	}

	for _, tt := range cases {
		if tt.in.IsNil() != tt.out {
			t.Errorf("%s.IsNil() shoud be %v", tt.in, tt.out)
		}
	}
}

func TestBoolIsAtom(t *testing.T) {
	cases := []struct {
		in  SExpr
		out bool
	}{
		{False, true},
		{True, true},
	}

	for _, tt := range cases {
		if tt.in.IsAtom() != tt.out {
			t.Errorf("%s.IsAtom() shoud be %v", tt.in, tt.out)
		}
	}
}

func TestBoolIsList(t *testing.T) {
	cases := []struct {
		in  SExpr
		out bool
	}{
		{False, false},
		{True, false},
	}

	for _, tt := range cases {
		if tt.in.IsList() != tt.out {
			t.Errorf("%s.IsList() shoud be %v", tt.in, tt.out)
		}
	}
}

func TestBoolIsCons(t *testing.T) {
	cases := []struct {
		in  SExpr
		out bool
	}{
		{False, false},
		{True, false},
	}

	for _, tt := range cases {
		if tt.in.IsCons() != tt.out {
			t.Errorf("%s.IsCons() shoud be %v", tt.in, tt.out)
		}
	}
}

func TestBoolString(t *testing.T) {
	cases := []struct {
		in  Bool
		out string
	}{
		{False, "#f"},
		{True, "#t"},
	}

	for _, tt := range cases {
		a := tt.in.String()
		if a != tt.out {
			t.Errorf("%v.String() == %q, want %q", tt.in, a, tt.out)
		}
	}
}
