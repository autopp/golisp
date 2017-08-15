package golisp

import (
	"testing"
)

func TestEqWithDifferentTypes(t *testing.T) {
	b := False
	n := Number(42)
	s := Symbol("foo")
	c := NewCons(Number(1), Number(2))
	cases := []struct {
		x, y SExpr
	}{
		{b, n},
		{b, s},
		{b, GetNil()},
		{b, c},
		{n, s},
		{n, GetNil()},
		{n, c},
		{s, GetNil()},
		{s, c},
		{GetNil(), c},
	}

	for _, tt := range cases {
		if Eq(tt.x, tt.y) {
			t.Errorf("Eq(%s, %s) == true, want false", tt.x, tt.y)
		}

		if Eq(tt.y, tt.x) {
			t.Errorf("Eq(%s, %s) == true, want false", tt.y, tt.x)
		}
	}
}

func TestEqWithSameTypes(t *testing.T) {
	cases := []struct {
		x, y SExpr
		out  bool
	}{
		{False, False, true},
		{False, True, false},
		{True, False, false},
		{True, True, true},
		{Number(42), Number(666), false},
		{Number(42), Number(42), true},
		{Symbol("foo"), Symbol("bar"), false},
		{Symbol("foo"), Symbol("foo"), true},
		{GetNil(), GetNil(), true},
		{NewCons(False, True), NewCons(True, True), false},
		{NewCons(False, True), NewCons(False, False), false},
		{NewCons(False, True), NewCons(False, True), true},
		{NewCons(NewCons(False, True), NewCons(True, False)), NewCons(NewCons(False, False), NewCons(True, False)), false},
		{NewCons(NewCons(False, True), NewCons(True, False)), NewCons(NewCons(False, True), NewCons(True, True)), false},
		{NewCons(NewCons(False, True), NewCons(True, False)), NewCons(NewCons(False, True), NewCons(True, False)), true},
	}

	for _, tt := range cases {
		got := Eq(tt.x, tt.y)
		if got != tt.out {
			t.Errorf("Eq(%s, %s) == %v, want %v", tt.x, tt.y, got, tt.out)
		}
	}
}
