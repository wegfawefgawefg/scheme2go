package parse

import (
	"fmt"
	"scheme2go/pkg/tkn"
	"strings"
)

// const (
// 	NumberLiteral  = "NumberLiteral"
// 	StringLiteral  = "StringLiteral"
// 	CallExpression = "CallExpression"
// 	Program        = "Program"
// )

type ParseResult struct {
	NextPosition int
	Node         Node
}

type NodeType int

const (
	NodeTypeNumber NodeType = iota
	NodeTypeString
	NodeTypeExpression
	NodeTypePackage
	NodeTypeProgn
	NodeTypeSymbol
	NodeTypeType
	NodeTypeFunctionDef
	NodeTypeFunctionParams
	NodeTypeFunctionReturnTypes
	NodeTypeFunctionBody
	// NodeTypeImport
	// NodeTypeImports
)

func (t NodeType) String() string {
	return [...]string{
		"Number",
		"String",
		"Expression",
		"Package",
		"Progn",
		"Symbol",
		"Type",
		"FunctionDef",
		"FunctionParams",
		"FunctionReturnTypes",
		"FunctionBody",
	}[t]
}

type NodeMetadata struct {
	File string
	Line int
	Col  int
}

type Node struct {
	Typ    NodeType
	Value  string
	Params []Node
}

func PrintNode(n Node, indent int) string {
	var result strings.Builder
	indentStr := strings.Repeat("  ", indent)
	result.WriteString(fmt.Sprintf("%sNode{\n", indentStr))
	result.WriteString(fmt.Sprintf("%s  Typ:   \"%s\",\n", indentStr, n.Typ.String()))
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

func ParsePackage(tokens []tkn.Token, current int) (ParseResult, error) {
	// (package main)
	return ParseResult{
		NextPosition: current + 2,
		Node: Node{
			Typ:    NodeTypePackage,
			Value:  tokens[current+1].Value,
			Params: []Node{},
		},
	}, nil
}

func ParseImport(tokens []tkn.Token, current int) (ParseResult, error) {
	// (import "fmt")
	return ParseResult{
		NextPosition: current + 2,
		Node: Node{
			Typ:    NodeTypePackage,
			Value:  tokens[current+1].Value,
			Params: []Node{},
		},
	}, nil
}

func ParseProgn(tokens []tkn.Token, current int) (ParseResult, error) {
	/*
		(progn
			(+ 1 2)
			(fmt.Println "hello")
			4
		)
	*/
	node := Node{
		Typ:    NodeTypeProgn,
		Value:  "",
		Params: []Node{},
	}

	current += 1 // skip the 'progn' keyword
	if current >= len(tokens) {
		return ParseResult{}, fmt.Errorf("unexpected end of tokens after 'progn'")
	}

	for tokens[current].Typ != tkn.TokenTypeParen || tokens[current].Value != ")" {
		if current >= len(tokens) {
			return ParseResult{}, fmt.Errorf("missing closing parenthesis for 'progn'")
		}
		result, err := ParseToken(tokens, current)
		if err != nil {
			return ParseResult{}, err
		}
		node.Params = append(node.Params, result.Node)
		current = result.NextPosition
	}

	current += 1 // skip closing parenthesis
	return ParseResult{
		NextPosition: current,
		Node:         node,
	}, nil
}

func ParseFunctionParams(tokens []tkn.Token, current int) ([]Node, int, error) {
	/*
		((a int) (b int))

		becomes

		Node{
			Typ: NodeTypeFunctionParams,
			Value: "",
			Params: [
				Node{
					Typ: NodeTypeSymbol,
					Value: "a",
				},
				Node{
					Typ: NodeTypeType,
					Value: "int",
				},
				Node{
					Typ: NodeTypeSymbol,
					Value: "b",
				},
				Node{
					Typ: NodeTypeType,
					Value: "int",
				},
			],
		}
	*/
	if tokens[current].Typ != tkn.TokenTypeParen || tokens[current].Value != "(" {
		return nil, current, fmt.Errorf("expected opening parenthesis for function parameters, got %v", tokens[current].Value)
	}
	current += 1 // Skip the opening parenthesis

	params := []Node{}
	for tokens[current].Typ != tkn.TokenTypeParen || tokens[current].Value != ")" {
		if tokens[current].Typ != tkn.TokenTypeParen || tokens[current].Value != "(" {
			return nil, current, fmt.Errorf("expected opening parenthesis for a parameter, got %v", tokens[current].Value)
		}
		current += 1 // Skip the opening parenthesis for a single parameter

		if tokens[current].Typ != tkn.TokenTypeSymbol {
			return nil, current, fmt.Errorf("expected parameter name, got %v", tokens[current].Value)
		}
		paramName := tokens[current].Value
		current += 1 // Move past parameter name

		if tokens[current].Typ != tkn.TokenTypeSymbol {
			return nil, current, fmt.Errorf("expected parameter type, got %v", tokens[current].Value)
		}
		paramType := tokens[current].Value
		current += 1 // Move past parameter type

		if tokens[current].Typ != tkn.TokenTypeParen || tokens[current].Value != ")" {
			return nil, current, fmt.Errorf("expected closing parenthesis for a parameter, got %v", tokens[current].Value)
		}
		current += 1 // Skip the closing parenthesis for a single parameter

		params = append(params, Node{
			Typ:   NodeTypeSymbol,
			Value: paramName,
			Params: []Node{
				{
					Typ:   NodeTypeType,
					Value: paramType,
				},
			},
		})

		// Check if we're at the end of the parameter list
		if tokens[current].Typ == tkn.TokenTypeParen && tokens[current].Value == ")" {
			break
		}
	}
	current += 1 // Skip the closing parenthesis of the parameters list

	return params, current, nil
}

func ParseFunctionReturnTypes(tokens []tkn.Token, current int) ([]Node, int, error) {
	/*
		(int int)

		becomes

		Node{
			Typ: NodeTypeFunctionReturnTypes,
			Value: "",
			Params: [
				Node{
					Typ: NodeTypeType,
					Value: "int",
				},
				Node{
					Typ: NodeTypeType,
					Value: "int",
				},
			],
		}

	*/
	if tokens[current].Typ != tkn.TokenTypeParen || tokens[current].Value != "(" {
		return nil, current, fmt.Errorf("expected opening parenthesis for return types, got %v", tokens[current].Value)
	}
	current += 1 // Skip the opening parenthesis

	returnTypes := []Node{}
	// Loop until we reach the closing parenthesis
	for tokens[current].Typ != tkn.TokenTypeParen || tokens[current].Value != ")" {
		if tokens[current].Typ != tkn.TokenTypeSymbol {
			return nil, current, fmt.Errorf("expected type, got %v", tokens[current])
		}

		// Create a NodeTypeType node for each return type
		returnType := Node{
			Typ:   NodeTypeType,
			Value: tokens[current].Value,
		}
		returnTypes = append(returnTypes, returnType)
		current += 1

		// If there are multiple return types, they will be separated by spaces
		// so we continue parsing until the closing parenthesis
	}

	if tokens[current].Typ != tkn.TokenTypeParen || tokens[current].Value != ")" {
		return nil, current, fmt.Errorf("expected closing parenthesis for return types, got %v", tokens[current].Value)
	}
	current += 1 // Skip the closing parenthesis

	return returnTypes, current, nil
}

// This is for the body of a function, which can contain multiple expressions
func ParseBody(tokens []tkn.Token, current int) (Node, int, error) {
	bodyNode := Node{
		Typ:    NodeTypeFunctionBody,
		Params: []Node{},
	}

	// Loop to process all statements within the body
	for tokens[current].Typ != tkn.TokenTypeParen || tokens[current].Value != ")" {
		if current >= len(tokens) {
			return Node{}, current, fmt.Errorf("missing closing parenthesis in function body")
		}
		result, err := ParseToken(tokens, current)
		if err != nil {
			return Node{}, current, err
		}
		bodyNode.Params = append(bodyNode.Params, result.Node)
		current = result.NextPosition
	}

	return bodyNode, current, nil
}

func ParseDefun(tokens []tkn.Token, current int) (ParseResult, error) {
	/*
		// nothing example
		(defun nothing () () ())

		// simple example
		(defun OneThing ((a int)) (int)
			(+ a 1)
		)

		// full function form
		(defun TwoThings ((a int) (b int)) (int int)
			(tuple int (+ a 1) (+ b 1))
		)
	*/
	current += 1 // Skip 'defun' keyword

	if tokens[current].Typ != tkn.TokenTypeSymbol {
		return ParseResult{}, fmt.Errorf("expected function name, got %v", tokens[current].Value)
	}

	functionName := tokens[current].Value
	current += 1 // Move past function name

	// Parse parameters
	functionParams, newCurrent, err := ParseFunctionParams(tokens, current)
	if err != nil {
		return ParseResult{}, err
	}
	current = newCurrent

	// Parse return types
	returnTypes, newCurrent, err := ParseFunctionReturnTypes(tokens, current)
	if err != nil {
		return ParseResult{}, err
	}
	current = newCurrent

	// Parse function body
	body, newCurrent, err := ParseBody(tokens, current)
	if err != nil {
		return ParseResult{}, err
	}
	current = newCurrent

	// Combine all parts into the function definition node
	params := []Node{}
	params = append(params, functionParams...)
	params = append(params, returnTypes...)
	params = append(params, body)

	functionDefinitionNode := Node{
		Typ:    NodeTypeFunctionDef,
		Value:  functionName,
		Params: params,
	}

	return ParseResult{
		NextPosition: current,
		Node:         functionDefinitionNode,
	}, nil
}

func ParseNumber(tokens []tkn.Token, current int) (ParseResult, error) {
	return ParseResult{
		NextPosition: current + 1,
		Node: Node{
			Typ:    NodeTypeNumber,
			Value:  tokens[current].Value,
			Params: []Node{},
		},
	}, nil
}

func ParseString(tokens []tkn.Token, current int) (ParseResult, error) {
	return ParseResult{
		NextPosition: current + 1,
		Node: Node{
			Typ:    NodeTypeString,
			Value:  tokens[current].Value,
			Params: []Node{},
		},
	}, nil
}

func ParseExpression(tokens []tkn.Token, current int) (ParseResult, error) {
	current += 1 // skip opening parenthesis

	token := tokens[current]
	switch token.Value {
	case "progn":
		return ParseProgn(tokens, current)
	case "defun":
		return ParseDefun(tokens, current)
	// Add other special forms as necessary
	default:
		node := Node{
			Typ:    NodeTypeExpression,
			Value:  token.Value,
			Params: []Node{},
		}

		current += 1 // move to the first argument
		for !(tokens[current].Typ == tkn.TokenTypeParen && tokens[current].Value == ")") {
			result, err := ParseToken(tokens, current)
			if err != nil {
				return ParseResult{}, err
			}
			node.Params = append(node.Params, result.Node)
			current = result.NextPosition
		}

		current += 1 // skip closing parenthesis
		return ParseResult{
			NextPosition: current,
			Node:         node,
		}, nil
	}
}

func ParseToken(tokens []tkn.Token, current int) (ParseResult, error) {
	token := tokens[current]

	switch token.Typ {
	case tkn.TokenTypeNumber:
		return ParseNumber(tokens, current)
	case tkn.TokenTypeString:
		return ParseString(tokens, current)
	case tkn.TokenTypeParen:
		if token.Value == "(" {
			return ParseExpression(tokens, current)
		}
		fallthrough
	default:
		return ParseResult{}, fmt.Errorf("unexpected token: %v", token)
	}
}

func ParseProgram(tokens []tkn.Token) (Node, error) {
	current := 0
	node := Node{
		Typ:    NodeTypeProgn,
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
