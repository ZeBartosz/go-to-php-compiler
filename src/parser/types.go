package parser

import (
	"fmt"

	"github.com/ZeBartosz/go-to-php-compiler/src/ast"
	"github.com/ZeBartosz/go-to-php-compiler/src/lexer"
)

type type_nud_handler func(p *parser) ast.Type
type type_led_handler func(p *parser, left ast.Type, bp binding_power) ast.Type

type type_nud_lookup map[lexer.TokenKind]type_nud_handler
type type_led_lookup map[lexer.TokenKind]type_led_handler
type type_bp_lookup map[lexer.TokenKind]binding_power

var type_bp_lu = type_bp_lookup{}
var type_nud_lu = type_nud_lookup{}
var type_led_lu = type_led_lookup{}

func type_led(kind lexer.TokenKind, bp binding_power, led_fn type_led_handler) {
	type_bp_lu[kind] = bp
	type_led_lu[kind] = led_fn
}

func type_nud(kind lexer.TokenKind, nud_fn type_nud_handler) {
	type_nud_lu[kind] = nud_fn
}

func createTokenTypeLookups() {
	type_nud(lexer.IDENTIFIER, parse_symbol_type)
	type_nud(lexer.OPEN_BRACKET, parse_array_type)
}

func parse_symbol_type(p *parser) ast.Type {
	return ast.SymbolType{
		Name: p.expect(lexer.IDENTIFIER).Value,
	}
}

func parse_array_type(p *parser) ast.Type {
	p.advance()
	p.expect(lexer.CLOSE_BRACKET)

	var UnderlyingType = parse_type(p, defalt_bp)
	return ast.ArrayType{
		Underlying: UnderlyingType,
	}
}

func parse_type(p *parser, bp binding_power) ast.Type {
	// gets the current token kind
	tokenKind := p.currentTokenKind()

	// checks nud look up for the token kind
	nud_fn, exists := type_nud_lu[tokenKind]

	// if it doesnt exists through an error
	if !exists {
		panic(fmt.Sprintf("TYPE_NUD HANDLER EXPECTED FOR TOKEN %s\n", lexer.TokenKindString(tokenKind)))
	}

	left := nud_fn(p)
	// checks if the current tokenKind has higher binding power
	for type_bp_lu[p.currentTokenKind()] > bp {
		// next tokenKind
		tokenKind := p.currentTokenKind()
		// checks if it exists in the look up
		led_fn, exists := type_led_lu[tokenKind]

		// if it doesnt provide an error
		if !exists {
			panic(fmt.Sprintf("TYPE_LED HANDLER EXPECTED FOR TOKEN %s\n", lexer.TokenKindString(tokenKind)))
		}

		left = led_fn(p, left, type_bp_lu[tokenKind])
	}

	return left
}
