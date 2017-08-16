package golisp

import (
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
