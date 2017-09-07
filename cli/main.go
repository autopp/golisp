package main

import (
	"fmt"
	"io"
	"os"

	"github.com/autopp/golisp"
)

func main() {
	var source io.Reader
	var filename string
	var err error

	switch len(os.Args) {
	case 1:
		source = os.Stdin
		filename = "<stdin>"
	case 2:
		source, err = os.Open(os.Args[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "cannot open %s\n", os.Args[1])
		}
		filename = os.Args[1]
	default:
		fmt.Fprintf(os.Stderr, "usage: %s [filename]\n", os.Args[0])
		os.Exit(1)
	}

	sexprs, err := golisp.Parse(source, filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot parse %s: %v\n", filename, err)
		os.Exit(1)
	}

	env := golisp.NewGlobalEnv()
	for _, sexpr := range sexprs {
		_, err := golisp.EvalSExpr(sexpr, env)
		if err != nil {
			fmt.Fprintf(os.Stderr, "runtime error: %s\n", err.Error())
			os.Exit(1)
		}
	}
}
