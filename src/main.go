package main

import (
	"os"

	"github.com/ZeBartosz/PrattParsing/src/lexer"
)

func main() {
	// we reading the file
	bytes, _ := os.ReadFile("./examples/01.lang")
	tokens := lexer.Tokenize(string(bytes))

	for _, token := range tokens {
		token.Debug()
	}
}
