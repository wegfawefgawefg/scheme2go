package parse

import (
	"fmt"
	"scheme2go/pkg/tokenize"
)

type ParseResult struct {
	NextPosition int
	Node         Node
}

type Node struct {
	Typ    string
	Value  string
	Params []Node
}

func ParseNumber(tokens []tokenize.Token, current int) (ParseResult, error) {
	return ParseResult{
		NextPosition: current + 1,
		Node: Node{
			Typ:    "NumberLiteral",
			Value:  tokens[current].Value,
			Params: []Node{},
		},
	}, nil
}

func ParseString(tokens []tokenize.Token, current int) (ParseResult, error) {
	return ParseResult{
		NextPosition: current + 1,
		Node: Node{
			Typ:    "StringLiteral",
			Value:  tokens[current].Value,
			Params: []Node{},
		},
	}, nil
}

func ParseExpression(tokens []tokenize.Token, current int) (ParseResult, error) {
	// steps:
	// skip opening parens
	// create base node with type CallExpression, and name from current token
	// recursively call parseToken until encountering a closing parens
	// skip the last token - the closing parens

	// skip opening parens
	current++
	token := tokens[current]
	node := Node{
		Typ:    "CallExpression",
		Value:  token.Value,
		Params: []Node{},
	}
	token = tokens[current+1]

	for !(token.Typ == "paren" && token.Value == ")") {
		// recursively call parseToken
		result, err := parseToken(tokens, current)
		if err != nil {
			return ParseResult{}, err
		}
		node.Params = append(node.Params, result.Node)
		current = result.NextPosition
		token = tokens[current]
	}

	current++
	return ParseResult{
		NextPosition: current,
		Node:         node,
	}, nil
}

func parseToken(tokens []tokenize.Token, current int) (ParseResult, error) {
	token := tokens[current]

	switch token.Typ {
	case "number":
		return ParseNumber(tokens, current)
	case "string":
		return ParseString(tokens, current)
	case "paren":
		if token.Value == "(" {
			return ParseExpression(tokens, current)
		}
		fallthrough
	default:
		return ParseResult{}, fmt.Errorf("unknown token type: %s", token.Typ)
	}
}
