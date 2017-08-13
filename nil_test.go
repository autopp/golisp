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
