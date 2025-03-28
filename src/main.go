package main

import (
	"fmt"
	"os"

	codegen "github.com/ZeBartosz/go-to-php-compiler/src/CodeGen"
	"github.com/ZeBartosz/go-to-php-compiler/src/lexer"
	"github.com/ZeBartosz/go-to-php-compiler/src/parser"
	"github.com/sanity-io/litter"
)

func main() {
	filePath := "./examples/05.lang"
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}
	tokens := lexer.Tokenize(string(bytes))

	ast, err := parser.Parse(tokens)
	if err != nil {
		fmt.Printf("Error parsing: %v\n", err)
		return
	}

	fmt.Println("--- Abstract Syntax Tree ---")
	litter.Dump(ast)

	// Generate PHP code from the AST
	fmt.Println("\n--- Generated PHP Code ---")

	generatedCode := codegen.GeneratePHP(ast)
	fmt.Println(generatedCode)
}
