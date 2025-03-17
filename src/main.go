package main

import (
	"fmt"
	"os"
)

func main() {
	// we reading the file
	bytes, _ := os.ReadFile("./examples/00.lang")
	source := string(bytes)

	fmt.Printf("code: %s\n", source)
}
