package golisp

import (
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	cases := []struct {
		source, name string
		out          SExpr
		err          bool
	}{
		{"#f", "false", False, false},
		{"#t", "true", True, false},
		{"42", "number", Number(42), false},
		{"+42", "pos_number", Number(42), false},
		{"-42", "neg_number", Number(-42), false},
		{"let-rec", "symbol", Symbol("let-rec"), false},
		{"( )", "nil", GetNil(), false},
		{"(#t . #f)", "pair", NewCons(True, False), false},
		{"(1 2 3)", "list", MakeList(Number(1), Number(2), Number(3)), false},
		{"() ,", "extra", GetNil(), true},
	}

	for _, tt := range cases {
		filename := tt.name + ".lsp"
		got, err := Parse(strings.NewReader(tt.source), filename)
		if tt.err {
			if err == nil {
				t.Errorf("Parse(%q, %q) == (%s, nil), want error", tt.source, filename, got)
			}
		} else {
			if err != nil {
				t.Errorf("Parse(%q, %q) == (%s, %s), want (%s, nil)", tt.source, filename, got, err, tt.out)
			} else if !Eq(got[0], tt.out) {
				t.Errorf("Parse(%q, %q) == (%s, nil), want (%s, nil)", tt.source, filename, got, tt.out)
			}
		}
	}
}
