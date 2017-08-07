package golisp

import (
	"testing"
)

func TestParseWithFalse(t *testing.T) {
	res, err := Parse("#f", "false.lsp")

	if err != nil {
		t.Fatal("Parse with #f shoud not return error, but got:", err.Error())
	}

	if res != False {
		t.Fatal("Parse with #f shoud return false value")
	}
}

func TestParseWithTrue(t *testing.T) {
	res, err := Parse("#t", "true.lsp")

	if err != nil {
		t.Fatal("Parse with #t shoud not return error, but got:", err.Error())
	}

	if res != True {
		t.Fatal("Parse with #t shoud return true value")
	}
}

func TestParseWithNumber(t *testing.T) {
	res, err := Parse("42", "number.lsp")

	if err != nil {
		t.Fatal("Parse with 42 shoud not return error, but got:", err.Error())
	}

	if res != Number(42) {
		t.Fatal("Parse with 42 shoud return number value")
	}
}

func TestParseWithPosNumber(t *testing.T) {
	res, err := Parse("+42", "pos_number.lsp")

	if err != nil {
		t.Fatal("Parse with +42 shoud not return error, but got:", err.Error())
	}

	if res != Number(42) {
		t.Fatal("Parse with +42 shoud return number value")
	}
}

func TestParseWithNegNumber(t *testing.T) {
	res, err := Parse("-42", "neg_number.lsp")

	if err != nil {
		t.Fatal("Parse with -42 shoud not return error, but got:", err.Error())
	}

	if res != Number(-42) {
		t.Fatal("Parse with -42 shoud return number value")
	}
}

func TestParseWithSymbol(t *testing.T) {
	res, err := Parse("let-rec", "symbol.lsp")

	if err != nil {
		t.Fatal("Parse with let-rec shoud not return error, but got:", err.Error())
	}

	if res != Symbol("let-rec") {
		t.Fatal("Parse with let-rec shoud return number value")
	}
}

func TestParseWithNil(t *testing.T) {
	res, err := Parse("()", "nil.lsp")

	if err != nil {
		t.Fatal("Parse with () shoud not return error, but got:", err.Error())
	}

	if !res.IsNil() {
		t.Fatal("Parse with () shoud return nil value")
	}
}
