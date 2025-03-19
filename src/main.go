package main

import (
	"os"

	"github.com/ZeBartosz/PrattParsing/src/lexer"
	"github.com/ZeBartosz/PrattParsing/src/parser"
	"github.com/sanity-io/litter"
)

func main() {
	// we reading the file
	bytes, _ := os.ReadFile("./examples/03.lang")
	tokens := lexer.Tokenize(string(bytes))
	ast := parser.Parse(tokens)

	litter.Dump(ast)
}
