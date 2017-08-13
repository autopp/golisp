package golisp

import (
	"testing"
)

func TestBoolIsNil(t *testing.T) {
	if False.IsNil() || True.IsNil() {
		t.Fatal("Bool.IsNil() should be false")
	}
}

func TestBoolIsAtom(t *testing.T) {
	if !False.IsAtom() || !True.IsAtom() {
		t.Fatal("Bool.IsAtom() should be true")
	}
}

func TestBoolIsList(t *testing.T) {
	if False.IsList() || True.IsList() {
		t.Fatal("Bool.IsList() should be false")
	}
}

func TestBoolIsCons(t *testing.T) {
	if False.IsCons() || True.IsCons() {
		t.Fatal("Bool.IsCons() should be false")
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
