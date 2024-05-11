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

func EmitNode(node parse.Node) (string, error) {
	switch node.Typ {
	case "Program":
		return EmitProgram(node), nil
	case "CallExpression":
		return EmitExpression(node)
	case "NumberLiteral":
		return EmitNumber(node), nil
	case "StringLiteral":
		return EmitString(node), nil
	default:
		return "", fmt.Errorf("unknown node type: %s", node.Typ)
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
