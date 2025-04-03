package ast

import "fmt"

func TypeToString(t Type) string {
	switch v := t.(type) {
	case SymbolType:
		return v.Name
	// Add additional cases
	default:
		return fmt.Sprintf("%v", t)
	}
}
