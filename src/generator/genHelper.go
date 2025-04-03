package generator

import "strings"

func (gen *Generator) indent() string {
	return strings.Repeat("    ", gen.indentLevel) // Four spaces per indent level.
}

func (gen *Generator) writeln(line string) {
	gen.strBuilder.WriteString(gen.indent())
	gen.strBuilder.WriteString(line)
	gen.strBuilder.WriteString("\n")
}

func (gen *Generator) increaseIndent() {
	gen.indentLevel++
}

func (gen *Generator) decreaseIndent() {
	if gen.indentLevel > 0 {
		gen.indentLevel--
	}
}
