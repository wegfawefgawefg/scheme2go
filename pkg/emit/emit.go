package emit

import (
	"fmt"
	"scheme2go/pkg/parse"
	"strings"
)

func EmitNumber(node parse.Node) string {
	return node.Value
}

func EmitString(node parse.Node) string {
	return fmt.Sprintf("\"%s\"", node.Value)
}

func EmitExpression(node parse.Node) (string, error) {
	// (symbol param param) => symbol(param, param)
	var result []string
	if len(node.Params) == 0 {
		return "", fmt.Errorf("expected at least one parameter")
	}

	result = append(result, node.Value)
	for _, param := range node.Params {
		e, err := EmitNode(param)
		if err != nil {
			return "", err
		}
		result = append(result, e)
	}
	return fmt.Sprintf("%s(%s)", result[0], strings.Join(result[1:], ", ")), nil
}

func EmitPackage(node parse.Node) string {
	// Example: (package main) => package main
	return fmt.Sprintf("package %s", node.Value)
}

func EmitProgn(node parse.Node) (string, error) {
	var result strings.Builder
	for _, param := range node.Params {
		e, err := EmitNode(param)
		if err != nil {
			return "", err
		}
		result.WriteString(e)
		result.WriteString("\n")
	}
	return result.String(), nil
}

func EmitSymbol(node parse.Node) string {
	// Directly return the symbol's name
	return node.Value
}

func EmitType(node parse.Node) string {
	// Directly return the type name
	return node.Value
}

func EmitFunctionDef(node parse.Node) (string, error) {
	if len(node.Params) < 3 {
		return "", fmt.Errorf("function definition requires at least parameters, return types, and body")
	}
	params, err := EmitFunctionParams(node.Params[0])
	if err != nil {
		return "", err
	}
	returnTypes, err := EmitFunctionReturnTypes(node.Params[1])
	if err != nil {
		return "", err
	}
	body, err := EmitFunctionBody(node.Params[2])
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("func %s(%s) %s {\n%s}\n", node.Value, params, returnTypes, body), nil
}

func EmitFunctionParams(node parse.Node) (string, error) {
	/*
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

	var params []string
	for _, param := range node.Params {
		params = append(params, param.Value)
	}

	return strings.Join(params, ", "), nil
}

func EmitFunctionReturnTypes(node parse.Node) (string, error) {
	// type, type, type
	var types []string
	for _, t := range node.Params {
		types = append(types, t.Value)
	}
	if len(types) == 1 {
		return types[0], nil
	}
	return fmt.Sprintf("(%s)", strings.Join(types, ", ")), nil
}

func EmitFunctionBody(node parse.Node) (string, error) {
	body, err := EmitProgn(node)
	if err != nil {
		return "", err
	}
	return body, nil
}

func EmitNode(node parse.Node) (string, error) {
	switch node.Typ {
	case parse.NodeTypeNumber:
		return EmitNumber(node), nil
	case parse.NodeTypeString:
		return EmitString(node), nil
	case parse.NodeTypeExpression:
		return EmitExpression(node)
	case parse.NodeTypePackage:
		return EmitPackage(node), nil
	case parse.NodeTypeProgn:
		return EmitProgn(node)
	case parse.NodeTypeSymbol:
		return EmitSymbol(node), nil
	case parse.NodeTypeType:
		return EmitType(node), nil
	case parse.NodeTypeFunctionDef:
		return EmitFunctionDef(node)
	case parse.NodeTypeFunctionParams:
		return EmitFunctionParams(node)
	case parse.NodeTypeFunctionReturnTypes:
		return EmitFunctionReturnTypes(node)
	case parse.NodeTypeFunctionBody:
		return EmitFunctionBody(node)
	default:
		return "", fmt.Errorf("unknown node type: %v", node.Typ)
	}
}

func EmitProgram(node parse.Node) string {
	var result strings.Builder
	for _, n := range node.Params {
		e, _ := EmitNode(n)
		result.WriteString(e)
		result.WriteString("\n")
	}
	return result.String()
}
