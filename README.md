# Go-to-PHP Compiler

## Overview

The Go-to-PHP Compiler is a project designed to help me learn about lexer, tokenization, and code generation. The primary goal of this compiler is to translate Go code into PHP code, allowing for a better understanding of how compilers work and the intricacies of language translation.

## Features

- **Lexer and Tokenization**: The compiler includes a lexer that tokenizes Go source code into meaningful tokens, which are then used for parsing.
- **Parsing**: The project implements a parser that constructs an Abstract Syntax Tree (AST) from the tokenized input.
- **Code Generation**: The compiler generates PHP code from the AST, allowing for the execution of Go-like code in a PHP environment.

## Learning Objectives

This project serves as a hands-on learning experience for the following concepts:

- Understanding how lexers work and how they tokenize source code.
- Learning about parsing techniques and constructing an AST.
- Exploring code generation and how to translate one programming language into another.

## Special Thanks

A special thanks to [Tyler Laceby](**https://www.youtube.com/channel/UC1g1g0g0g0g0g0g0g0g0g0**) on YouTube for his excellent explanations of lexer and tokenization. His tutorials have been instrumental in guiding me through the process of building this compiler.

## Getting Started

To get started with the Go-to-PHP Compiler, clone the repository and run the following commands:

```bash
git clone https://github.com/ZeBartosz/go-to-php-compiler.git
cd go-to-php-compiler
go run main.go
```

## Contributing

If you have suggestions for improvements or would like to contribute to the project, feel free to open an issue or submit a pull request.

## License 

This project is licensed under the MIT License.
