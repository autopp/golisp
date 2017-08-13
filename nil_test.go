package golisp

import (
	"testing"
)

func TestNilIsNil(t *testing.T) {
	if !GetNil().IsNil() {
		t.Fatal("Nil.IsNil() should be true")
	}
}

func TestNilIsAtom(t *testing.T) {
	if !GetNil().IsAtom() {
		t.Fatal("Nil.IsAtom() should be true")
	}
}

func TestNilIsList(t *testing.T) {
	if !GetNil().IsList() {
		t.Fatal("Nil.IsList() should be true")
	}
}

func TestNilIsCons(t *testing.T) {
	if GetNil().IsCons() {
		t.Fatal("Nil.IsCons() should be false")
	}
}

func TestNilString(t *testing.T) {
	in := GetNil().(*nilImpl)
	out := "()"
	a := in.String()

	if in.String() != out {
		t.Errorf("GetNil().String() == %q, want %q", a, out)
	}
}
