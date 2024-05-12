package emit

import (
	"scheme2go/pkg/parse"
	"testing"
)

func TestEmitNumber(t *testing.T) {
	input := parse.Node{Typ: parse.NodeTypeNumber, Value: "12345", Params: []parse.Node{}}
	result := EmitNumber(input)
	expected := "12345"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestEmitString(t *testing.T) {
	input := parse.Node{Typ: parse.NodeTypeString, Value: "abc 123", Params: []parse.Node{}}
	result := EmitString(input)
	expected := "\"abc 123\""
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

// func TestEmitExpression(t *testing.T) {
// 	input := parse.Node{
// 		Typ:   "CallExpression",
// 		Value: "subtract",
// 		Params: []parse.Node{
// 			{Typ: "NumberLiteral", Value: "4", Params: []parse.Node{}},
// 			{Typ: "NumberLiteral", Value: "2", Params: []parse.Node{}},
// 		},
// 	}
// 	result, err := EmitExpression(input)
// 	if err != nil {
// 		t.Errorf("Unexpected error: %v", err)
// 	}
// 	expected := "subtract(4, 2)"
// 	if result != expected {
// 		t.Errorf("Expected %s, got %s", expected, result)
// 	}
// }

// func TestEmitProgram(t *testing.T) {
// 	input := parse.Node{
// 		Typ: "Program",
// 		Params: []parse.Node{
// 			// print "lalala"
// 			{
// 				Typ:    "CallExpression",
// 				Value:  "print",
// 				Params: []parse.Node{{Typ: "StringLiteral", Value: "lalala", Params: []parse.Node{}}},
// 			},
// 			// (add 2 (sub 4 2))
// 			{
// 				Typ:   "CallExpression",
// 				Value: "add",
// 				Params: []parse.Node{
// 					{Typ: "NumberLiteral", Value: "2", Params: []parse.Node{}},
// 					{
// 						Typ:   "CallExpression",
// 						Value: "sub",
// 						Params: []parse.Node{
// 							{Typ: "NumberLiteral", Value: "4", Params: []parse.Node{}},
// 							{Typ: "NumberLiteral", Value: "2", Params: []parse.Node{}},
// 						},
// 					},
// 				},
// 			},
// 		},
// 	}
// 	result := EmitProgram(input)
// 	expected := "print(\"lalala\")\nadd(2, sub(4, 2))\n"
// 	if result != expected {
// 		t.Errorf("Expected %s, got %s", expected, result)
// 	}
// }
