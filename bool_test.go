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

func TestBoolEqWithBool(t *testing.T) {
	if !True.Eq(True) {
		t.Fatal("Bool.Eq() with same should be true")
	}

	if !False.Eq(False) {
		t.Fatal("Bool.Eq() with same should be true")
	}

	if True.Eq(False) {
		t.Fatal("Bool.Eq() with different should be false")
	}

	if False.Eq(True) {
		t.Fatal("Bool.Eq() with different should be false")
	}
}
