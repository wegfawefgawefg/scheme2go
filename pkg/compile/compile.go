package compile

import (
	"fmt"
	"scheme2go/pkg/emit"
	"scheme2go/pkg/parse"
	"scheme2go/pkg/tkn"
)

func Compile(input string) (string, error) {
	tokens, err := tkn.Tokenize(input)
	if err != nil {
		return "", err
	}
	ast, err := parse.ParseProgram(tokens)
	if err != nil {
		return "", fmt.Errorf("error parsing tokens: %v", err)
	}
	output, err := emit.EmitNode(ast)
	if err != nil {
		return "", fmt.Errorf("error emitting AST: %v", err)
	}
	return output, nil
}
