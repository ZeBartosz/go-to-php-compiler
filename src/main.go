package main

import (
	"fmt"
	"os"

	"github.com/ZeBartosz/go-to-php-compiler/src/generator"
	"github.com/ZeBartosz/go-to-php-compiler/src/lexer"
	"github.com/ZeBartosz/go-to-php-compiler/src/parser"
)

func main() {
	filePath := "./examples/complex/06.go"
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

	// Debug
	// fmt.Println("--- Abstract Syntax Tree ---")
	// litter.Dump(ast)

	generatedCode := generator.GeneratePHP(ast)

	// Define the output file name
	outputFileName := "main.php"

	// Write the generated PHP code to the file
	err = os.WriteFile(outputFileName, []byte(generatedCode), 0644)
	if err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
		return
	}

	fmt.Printf("\n--- Generated PHP Code written to %s ---\n", outputFileName)
}
