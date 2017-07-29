package golisp

import (
  "testing"
)

func TestBoolIsNil(t *testing.T) {
  var b Bool
  if b.IsNil() {
    t.Fatal("Bool.IsNil() should be false")
  }
}

func TestBoolIsAtom(t *testing.T) {
  var b Bool
  if !b.IsAtom() {
    t.Fatal("Bool.IsAtom() should be true")
  }
}

func TestBoolIsNil(t *testing.T) {
  var b Bool
  if b.IsNil {
    t.Fatal("Bool.IsAtom() should be false")
  }
}

func TestNumberIsNil(t *testing.T) {
  var n Number
  if n.IsNil() {
    t.Fatal("Number.IsNil() should be false")
  }
}
