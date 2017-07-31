package sexpr

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

func TestNewCons(t *testing.T) {
  cons := NewCons(True, False)
  if cons.Car != True {
    t.Fatal("NewCons() shoud set 1st argument to car")
  }

  if cons.Cdr != False {
    t.Fatal("NewCons() shoud set 2nd argument to cdr")
  }
}

func TestConsIsNil(t *testing.T) {
	if NewCons(True, False).IsNil() {
		t.Fatal("Cons.IsNil() should be false")
	}
}

func TestConsIsAtom(t *testing.T) {
	if NewCons(True, False).IsAtom() {
		t.Fatal("Cons.IsAtom() should be true")
	}
}

func TestConsIsList(t *testing.T) {
  cons := NewCons(Number(1), NewCons(Number(2), Number(3)))
	if cons.IsList() {
		t.Fatal("When last cdr is not nil, Cons.IsList() should be false")
	}

  cons = NewCons(Number(1), NewCons(Number(2), GetNil()))
  if !cons.IsList() {
		t.Fatal("When last cdr is nil, Cons.IsList() should be true")
	}
}

func TestConsIsCons(t *testing.T) {
	if !NewCons(True, False).IsCons() {
		t.Fatal("Cons.IsCons() should be true")
	}
}
