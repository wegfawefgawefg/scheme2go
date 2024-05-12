package parse

import (
	"reflect"
	"scheme2go/pkg/tkn"
	"testing"
)

func TestParseNumber(t *testing.T) {
	input := []tkn.Token{{Typ: tkn.TokenTypeNumber, Value: "12345"}}
	result, err := ParseNumber(input, 0)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expected := ParseResult{
		NextPosition: 1,
		Node: Node{
			Typ:    NodeTypeNumber,
			Value:  "12345",
			Params: []Node{},
		},
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %+v, got %+v", expected, result)
	}
}

func TestParseString(t *testing.T) {
	input := []tkn.Token{{Typ: tkn.TokenTypeString, Value: "abc 123"}}
	result, err := ParseString(input, 0)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expected := ParseResult{
		NextPosition: 1,
		Node: Node{
			Typ:    NodeTypeString,
			Value:  "abc 123",
			Params: []Node{},
		},
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %+v, got %+v", expected, result)
	}
}

// func TestParseExpression(t *testing.T) {
// 	input := []tkn.Token{
// 		{Typ: tkn.TokenTypeParen, Value: "("},
// 		{Typ: tkn.TokenTypeSymbol, Value: "subtract"},
// 		{Typ: tkn.TokenTypeNumber, Value: "4"},
// 		{Typ: tkn.TokenTypeNumber, Value: "2"},
// 		{Typ: tkn.TokenTypeParen, Value: ")"},
// 	}
// 	result, err := ParseExpression(input, 0)
// 	if err != nil {
// 		t.Errorf("Unexpected error: %v", err)
// 	}
// 	expected := ParseResult{
// 		NextPosition: 5,
// 		Node: Node{
// 			Typ:   NodeTypeCallExpression,
// 			Value: "subtract",
// 			Params: []Node{
// 				{Typ: "NumberLiteral", Value: "4", Params: []Node{}},
// 				{Typ: "NumberLiteral", Value: "2", Params: []Node{}},
// 			},
// 		},
// 	}
// 	if !reflect.DeepEqual(result, expected) {
// 		ppExpected := PrintParseResult(expected)
// 		ppResult := PrintParseResult(result)
// 		t.Errorf("Expected \n%s\n, got \n%s", ppExpected, ppResult)
// 	}
// }

// func TestParseExpressionWithNestedExpression(t *testing.T) {
// 	input := []tkn.Token{
// 		{Typ: tkn.TokenTypeParen, Value: "("},
// 		{Typ: tkn.TokenTypeSymbol, Value: "add"},
// 		{Typ: tkn.TokenTypeNumber, Value: "2"},
// 		{Typ: tkn.TokenTypeParen, Value: "("},
// 		{Typ: tkn.TokenTypeSymbol, Value: "subtract"},
// 		{Typ: tkn.TokenTypeNumber, Value: "4"},
// 		{Typ: tkn.TokenTypeNumber, Value: "2"},
// 		{Typ: tkn.TokenTypeParen, Value: ")"},
// 		{Typ: tkn.TokenTypeParen, Value: ")"},
// 	}
// 	result, err := ParseExpression(input, 0)
// 	if err != nil {
// 		t.Errorf("Unexpected error: %v", err)
// 	}
// 	expected := ParseResult{
// 		NextPosition: 9,
// 		Node: Node{
// 			Typ:   "CallExpression",
// 			Value: "add",
// 			Params: []Node{
// 				{Typ: "NumberLiteral", Value: "2", Params: []Node{}},
// 				{
// 					Typ:   "CallExpression",
// 					Value: "subtract",
// 					Params: []Node{
// 						{Typ: "NumberLiteral", Value: "4", Params: []Node{}},
// 						{Typ: "NumberLiteral", Value: "2", Params: []Node{}},
// 					},
// 				},
// 			},
// 		},
// 	}
// 	if !reflect.DeepEqual(result, expected) {
// 		ppExpected := PrintParseResult(expected)
// 		ppResult := PrintParseResult(result)
// 		t.Errorf("Expected \n%s\n, got \n%s", ppExpected, ppResult)
// 	}

// 	// Test from the center
// 	result, err = ParseExpression(input, 3)
// 	if err != nil {
// 		t.Errorf("Unexpected error: %v", err)
// 	}
// 	expected = ParseResult{
// 		NextPosition: 8,
// 		Node: Node{
// 			Typ:   "CallExpression",
// 			Value: "subtract",
// 			Params: []Node{
// 				{Typ: "NumberLiteral", Value: "4", Params: []Node{}},
// 				{Typ: "NumberLiteral", Value: "2", Params: []Node{}},
// 			},
// 		},
// 	}
// 	if !reflect.DeepEqual(result, expected) {
// 		ppExpected := PrintParseResult(expected)
// 		ppResult := PrintParseResult(result)
// 		t.Errorf("Expected \n%s\n, got \n%s", ppExpected, ppResult)
// 	}
// }

// func TestParseProgram(t *testing.T) {
// 	input := []tkn.Token{
// 		{Typ: tkn.TokenTypeParen, Value: "("},
// 		{Typ: tkn.TokenTypeSymbol, Value: "print"},
// 		{Typ: tkn.TokenTypeString, Value: "Hello"},
// 		{Typ: tkn.TokenTypeNumber, Value: "2"},
// 		{Typ: tkn.TokenTypeParen, Value: ")"},
// 		{Typ: tkn.TokenTypeParen, Value: "("},
// 		{Typ: tkn.TokenTypeSymbol, Value: "add"},
// 		{Typ: tkn.TokenTypeNumber, Value: "2"},
// 		{Typ: tkn.TokenTypeParen, Value: "("},
// 		{Typ: tkn.TokenTypeSymbol, Value: "subtract"},
// 		{Typ: tkn.TokenTypeNumber, Value: "4"},
// 		{Typ: tkn.TokenTypeNumber, Value: "2"},
// 		{Typ: tkn.TokenTypeParen, Value: ")"},
// 		{Typ: tkn.TokenTypeParen, Value: ")"},
// 	}
// 	result, err := ParseProgram(input)
// 	if err != nil {
// 		t.Errorf("Unexpected error: %v", err)
// 	}
// 	expected := Node{
// 		Typ:   "Program",
// 		Value: "",
// 		Params: []Node{
// 			{
// 				Typ:   "CallExpression",
// 				Value: "print",
// 				Params: []Node{
// 					{Typ: "StringLiteral", Value: "Hello", Params: []Node{}},
// 					{Typ: "NumberLiteral", Value: "2", Params: []Node{}},
// 				},
// 			},
// 			{
// 				Typ:   "CallExpression",
// 				Value: "add",
// 				Params: []Node{
// 					{Typ: "NumberLiteral", Value: "2", Params: []Node{}},
// 					{
// 						Typ:   "CallExpression",
// 						Value: "subtract",
// 						Params: []Node{
// 							{Typ: "NumberLiteral", Value: "4", Params: []Node{}},
// 							{Typ: "NumberLiteral", Value: "2", Params: []Node{}},
// 						},
// 					},
// 				},
// 			},
// 		},
// 	}
// 	if !reflect.DeepEqual(result, expected) {
// 		t.Errorf("Expected %+v, got %+v", expected, result)
// 	}
// }
