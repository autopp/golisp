package golisp

import (
	"fmt"
	"testing"
)

func TestProcBaseMethods(t *testing.T) {
	name := "foo"
	req := 2
	op := 1
	b := &procBase{name, req, op}

	if got := b.Name(); got != name {
		t.Errorf("%v.Name() == %q, want %q", b, got, name)
	}

	if got := b.Required(); got != req {
		t.Errorf("%v.Required() == %d, want %d", b, got, req)
	}

	if got := b.Optional(); got != op {
		t.Errorf("%v.Optional() == %d, want %d", b, got, op)
	}
}

func TestSpFormString(t *testing.T) {
	n := "foo"
	in := NewSpForm(n, 0, 0, func(_ []SExpr, _ *Env) (SExpr, error) {
		return GetNil(), nil
	})
	out := fmt.Sprintf("#<special %s>", n)
	got := in.String()

	if got != out {
		t.Errorf("SpForm.String() == %q, want %q", got, out)
	}
}
