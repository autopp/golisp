package golisp

import (
	"testing"
)

func TestSymbolIsNil(t *testing.T) {
	var s Symbol
	if s.IsNil() {
		t.Fatal("Symbol.IsNil() should be false")
	}
}

func TestSymbolIsAtom(t *testing.T) {
	var s Symbol
	if !s.IsAtom() {
		t.Fatal("Symbol.IsAtom() should be true")
	}
}

func TestSymbolIsList(t *testing.T) {
	var s Symbol
	if s.IsList() {
		t.Fatal("Symbol.IsList() should be false")
	}
}

func TestSymbolIsCons(t *testing.T) {
	var s Symbol
	if s.IsCons() {
		t.Fatal("Symbol.IsCons() should be false")
	}
}

func TestSymbolIsProc(t *testing.T) {
	in := Symbol("foo")
	if in.IsProc() {
		t.Errorf("%s.IsProc == true, want false", in)
	}
}

func TestSymbolString(t *testing.T) {
	in := Symbol("foo")
	out := "foo"
	a := in.String()

	if a != out {
		t.Errorf("%s.String() == %q, want %q", "foo", a, out)
	}
}
