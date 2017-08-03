package golisp

import (
  "fmt"
  "strings"
  "bufio"
  "regexp"
  "errors"
)

type token struct {
  kind int
  line, col int
  source string
}

type tokenizeRule struct {
  pattern *regexp.Regexp
  action func (string) *token
}

func formatError(filename string, line, col int, message string) error {
  return errors.New(fmt.Sprintf("%s:%d:%d:%s", filename, line, col, message))
}

func Parse(source, filename string) (SExpr, error) {
  tokens, err := tokenize(bufio.NewReader(strings.NewReader(source)), filename)
  if err != nil {
    return GetNil(), err
  }
  sexpr, err := parseSExpr(tokens, filename)
  return sexpr, err
}

func tokenize(source *bufio.Reader, filename string) ([]token, error) {
  return nil, errors.New("not implemented")
}

func parseSExpr(tokens []token, filename string) (SExpr, error) {
  return GetNil(), errors.New("not implemented")
}
