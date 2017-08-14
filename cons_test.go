package golisp

import (
	"testing"
)

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

func TestConsString(t *testing.T) {
	cases := []struct {
		in  *Cons
		out string
	}{
		{NewCons(Number(1), GetNil()), "(1)"},
		{NewCons(Number(1), Number(2)), "(1 . 2)"},
		{NewCons(Number(1), NewCons(Number(2), Number(3))), "(1 2 . 3)"},
		{NewCons(Number(1), NewCons(Number(2), NewCons(Number(3), GetNil()))), "(1 2 3)"},
		{NewCons(NewCons(Number(1), NewCons(GetNil(), Number(2))), NewCons(Number(3), GetNil())), "((1 () . 2) 3)"},
	}

	for _, tt := range cases {
		a := tt.in.String()

		if a != tt.out {
			t.Errorf("%v.String() == %q, want %q", tt.in, a, tt.out)
		}
	}
}
