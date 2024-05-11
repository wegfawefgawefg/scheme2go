package parse

import (
	"fmt"
	"scheme2go/pkg/tokenize"
	"strings"
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

func PrintNode(n Node, indent int) string {
	var result strings.Builder
	indentStr := strings.Repeat("  ", indent)
	result.WriteString(fmt.Sprintf("%sNode{\n", indentStr))
	result.WriteString(fmt.Sprintf("%s  Typ:   \"%s\",\n", indentStr, n.Typ))
	result.WriteString(fmt.Sprintf("%s  Value: \"%s\",\n", indentStr, n.Value))
	if len(n.Params) > 0 {
		result.WriteString(fmt.Sprintf("%s  Params: [\n", indentStr))
		for _, param := range n.Params {
			result.WriteString(PrintNode(param, indent+2))
		}
		result.WriteString(fmt.Sprintf("%s  ],\n", indentStr))
	} else {
		result.WriteString(fmt.Sprintf("%s  Params: []Node{},\n", indentStr))
	}
	result.WriteString(fmt.Sprintf("%s},\n", indentStr))
	return result.String()
}

func PrintParseResult(pr ParseResult) string {
	var result strings.Builder
	result.WriteString("ParseResult{\n")
	result.WriteString(fmt.Sprintf("  NextPosition: %d,\n", pr.NextPosition))
	result.WriteString(PrintNode(pr.Node, 1))
	result.WriteString("}\n")
	return result.String()
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
	current += 1 // skip opening parenthesis
	token := tokens[current]
	node := Node{
		Typ:    "CallExpression",
		Value:  token.Value,
		Params: []Node{},
	}

	current += 1 // skip symbol
	token = tokens[current]

	for !(token.Typ == "paren" && token.Value == ")") {
		// recursively call parseToken
		result, err := ParseToken(tokens, current)
		if err != nil {
			return ParseResult{}, err
		}
		node.Params = append(node.Params, result.Node)
		current = result.NextPosition
		token = tokens[current]
	}

	current += 1 // skip closing parenthesis
	return ParseResult{
		NextPosition: current,
		Node:         node,
	}, nil
}

func ParseToken(tokens []tokenize.Token, current int) (ParseResult, error) {
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
		return ParseResult{}, fmt.Errorf("unexpected token: %v", token)
	}
}

func ParseProgram(tokens []tokenize.Token) (Node, error) {
	current := 0
	node := Node{
		Typ:    "Program",
		Value:  "",
		Params: []Node{},
	}

	for current < len(tokens) {
		result, err := ParseToken(tokens, current)
		if err != nil {
			return Node{}, err
		}
		node.Params = append(node.Params, result.Node)
		current = result.NextPosition
	}

	return node, nil
}
