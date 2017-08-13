package golisp

import (
	"testing"
)

func TestNumberIsNil(t *testing.T) {
	var n Number
	if n.IsNil() {
		t.Fatal("Number.IsNil() should be false")
	}
}

func TestNumberIsAtom(t *testing.T) {
	var n Number
	if !n.IsAtom() {
		t.Fatal("Number.IsAtom() should be true")
	}
}

func TestNumberIsList(t *testing.T) {
	var n Number
	if n.IsList() {
		t.Fatal("Number.IsList() should be false")
	}
}

func TestNumberIsCons(t *testing.T) {
	var n Number
	if n.IsCons() {
		t.Fatal("Number.IsCons() should be false")
	}
}

func TestNumberString(t *testing.T) {
	in := Number(42)
	out := "42"
	a := in.String()

	if a != out {
		t.Errorf("%v.String() == %q, want %q", in, a, out)
	}
}
