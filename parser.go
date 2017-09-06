package golisp

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"regexp"
	"strconv"
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
	filename  string
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

func (tk *token) String() string {
	return tk.source
}

func formatError(filename string, line, col int, format string, args ...interface{}) error {
	header := fmt.Sprintf("%s:%d:%d: ", filename, line, col)
	message := fmt.Sprintf(format, args...)
	return errors.New(header + message)
}

func formatErrorWithToken(token *token, format string, args ...interface{}) error {
	return formatError(token.filename, token.line, token.col, format, args...)
}

func Parse(source io.Reader, filename string) ([]SExpr, error) {
	tokens, err := tokenize(source, filename)
	if err != nil {
		return nil, err
	}
	ret := make([]SExpr, 0)
	for len(tokens) > 0 {
		var sexpr SExpr
		var err error
		sexpr, tokens, err = parseSExpr(tokens)
		if err != nil {
			return nil, err
		}
		ret = append(ret, sexpr)
	}

	return ret, err
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
						tokens = append(tokens, &token{tokenKind, filename, line, col, text[0:matched[1]]})
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
		line++
	}
	return tokens, nil
}

func parseSExpr(tokens []*token) (SExpr, []*token, error) {
	switch tokens[0].kind {
	case falseToken:
		return False, tokens[1:], nil
	case trueToken:
		return True, tokens[1:], nil
	case numberToken:
		numberValue, _ := strconv.Atoi(tokens[0].source)
		return Number(numberValue), tokens[1:], nil
	case symbolToken:
		return Symbol(tokens[0].source), tokens[1:], nil
	case lparenToken:
		return parseCons(tokens[1:])
	default:
		return GetNil(), tokens, formatErrorWithToken(tokens[0], "expected sexpr, but got `%s'", tokens[0])
	}
}

func parseCons(tokens []*token) (SExpr, []*token, error) {
	switch tokens[0].kind {
	case rparenToken:
		return GetNil(), tokens[1:], nil
	default:
		car, rest, err := parseSExpr(tokens)
		if err != nil {
			return GetNil(), rest, err
		}
		cdr, rest, err := parseConsRest(rest)
		if err != nil {
			return GetNil(), rest, err
		}
		return NewCons(car, cdr), rest, nil
	}
}

func parseConsRest(tokens []*token) (SExpr, []*token, error) {
	switch tokens[0].kind {
	case dotToken:
		cdr, rest, err := parseSExpr(tokens[1:])
		if err != nil {
			return GetNil(), rest, err
		}
		if rest[0].kind != rparenToken {
			return GetNil(), rest, formatErrorWithToken(rest[0], "expected `)', but got `%s'", rest[0])
		}
		return cdr, rest[1:], nil
	default:
		return GetNil(), tokens, formatErrorWithToken(tokens[0], "list literal is not implemented")
	}
}
