package golisp

import (
  "testing"
)

func TestParseWithNil(t *testing.T) {
  res, err := Parse("()", "")

  if err != nil {
    t.Fatal("Parse with () shoud not return error, but got:", err.Error())
  }

  if !res.IsNil() {
    t.Fatal("Parse with () shoud return nil value")
  }
}
