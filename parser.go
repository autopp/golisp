package golisp

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"regexp"
	"strings"
)

type tokenKind int

const (
	errToken tokenKind = iota - 1
	emptyToken
	falseToken
	trueToken
	numberToken
	symbolToken
	lparenToken
	rparenToken
	dotToken
	quoteToken
)

type token struct {
	kind      tokenKind
	line, col int
	source    string
}

type tokenizeRule struct {
	pattern *regexp.Regexp
	kind    tokenKind
}

var tokenizeRules = []tokenizeRule{
	tokenizeRule{regexp.MustCompile(`^[ \t\n]`), emptyToken},
	tokenizeRule{regexp.MustCompile(`^#f`), falseToken},
	tokenizeRule{regexp.MustCompile(`^#t`), trueToken},
	tokenizeRule{regexp.MustCompile(`^[-+]?(0|[1-9][0-9]*)`), numberToken},
	tokenizeRule{regexp.MustCompile(`^[-+*/!?_a-zA-Z][-+*/!?_a-zA-Z0-9]*`), symbolToken},
	tokenizeRule{regexp.MustCompile(`^\(`), lparenToken},
	tokenizeRule{regexp.MustCompile(`^\)`), rparenToken},
	tokenizeRule{regexp.MustCompile(`^\.`), dotToken},
	tokenizeRule{regexp.MustCompile(`^'`), quoteToken},
}

func formatError(filename string, line, col int, message string) error {
	return errors.New(fmt.Sprintf("%s:%d:%d:%s", filename, line, col, message))
}

func Parse(source, filename string) (SExpr, error) {
	tokens, err := tokenize(strings.NewReader(source), filename)
	if err != nil {
		return GetNil(), err
	}
	sexpr, err := parseSExpr(tokens, filename)
	return sexpr, err
}

func tokenize(source io.Reader, filename string) ([]*token, error) {
	tokens := make([]*token, 0)
	line := 1
	scanner := bufio.NewScanner(source)
	for scanner.Scan() {
		col := 1
		text := scanner.Text()
		for len(text) > 0 {
			tokenKind := errToken
			for _, rule := range tokenizeRules {
				if matched := rule.pattern.FindStringIndex(text); matched != nil {
					tokenKind = rule.kind
					if tokenKind != emptyToken {
						tokens = append(tokens, &token{tokenKind, line, col, text[0:matched[1]]})
					}
					col += matched[1]
					text = text[matched[1]:]
					break
				}
			}
			if tokenKind == errToken {
				return nil, formatError(filename, line, col, "unreconized charactor")
			}
		}
		line += 1
	}
	return tokens, nil
}

func parseSExpr(tokens []*token, filename string) (SExpr, error) {
	return GetNil(), errors.New("not implemented")
}
