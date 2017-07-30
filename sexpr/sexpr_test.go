package sexpr

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

func TestBoolIsList(t *testing.T) {
  var b Bool
  if b.IsList() {
    t.Fatal("Bool.IsList() should be false")
  }
}

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

func TestNilIsNil(t *testing.T) {
  var s Nil
  if !s.IsNil() {
    t.Fatal("Nil.IsNil() should be true")
  }
}

func TestNilIsAtom(t *testing.T) {
  var s Nil
  if !s.IsAtom() {
    t.Fatal("Nil.IsAtom() should be true")
  }
}

func TestNilIsList(t *testing.T) {
  var s Nil
  if !s.IsList() {
    t.Fatal("Nil.IsList() should be true")
  }
}
