package parse

import (
	"reflect"
	"scheme2go/pkg/tokenize"
	"testing"
)

func TestParseNumber(t *testing.T) {
	input := []tokenize.Token{{Typ: "number", Value: "12345"}}
	result, err := ParseNumber(input, 0)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expected := ParseResult{
		NextPosition: 1,
		Node: Node{
			Typ:    "NumberLiteral",
			Value:  "12345",
			Params: []Node{},
		},
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %+v, got %+v", expected, result)
	}
}

func TestParseString(t *testing.T) {
	input := []tokenize.Token{{Typ: "string", Value: "abc 123"}}
	result, err := ParseString(input, 0)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expected := ParseResult{
		NextPosition: 1,
		Node: Node{
			Typ:    "StringLiteral",
			Value:  "abc 123",
			Params: []Node{},
		},
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %+v, got %+v", expected, result)
	}
}

func TestParseExpression(t *testing.T) {
	input := []tokenize.Token{
		{Typ: "paren", Value: "("},
		{Typ: "symbol", Value: "subtract"},
		{Typ: "number", Value: "4"},
		{Typ: "number", Value: "2"},
		{Typ: "paren", Value: ")"},
	}
	result, err := ParseExpression(input, 0)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expected := ParseResult{
		NextPosition: 5,
		Node: Node{
			Typ:   "CallExpression",
			Value: "subtract",
			Params: []Node{
				{Typ: "NumberLiteral", Value: "4", Params: []Node{}},
				{Typ: "NumberLiteral", Value: "2", Params: []Node{}},
			},
		},
	}
	if !reflect.DeepEqual(result, expected) {
		ppExpected := PrintParseResult(expected)
		ppResult := PrintParseResult(result)
		t.Errorf("Expected \n%s\n, got \n%s", ppExpected, ppResult)
	}
}

func TestParseExpressionWithNestedExpression(t *testing.T) {
	input := []tokenize.Token{
		{Typ: "paren", Value: "("},
		{Typ: "symbol", Value: "add"},
		{Typ: "number", Value: "2"},
		{Typ: "paren", Value: "("},
		{Typ: "symbol", Value: "subtract"},
		{Typ: "number", Value: "4"},
		{Typ: "number", Value: "2"},
		{Typ: "paren", Value: ")"},
		{Typ: "paren", Value: ")"},
	}
	result, err := ParseExpression(input, 0)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expected := ParseResult{
		NextPosition: 9,
		Node: Node{
			Typ:   "CallExpression",
			Value: "add",
			Params: []Node{
				{Typ: "NumberLiteral", Value: "2", Params: []Node{}},
				{
					Typ:   "CallExpression",
					Value: "subtract",
					Params: []Node{
						{Typ: "NumberLiteral", Value: "4", Params: []Node{}},
						{Typ: "NumberLiteral", Value: "2", Params: []Node{}},
					},
				},
			},
		},
	}
	if !reflect.DeepEqual(result, expected) {
		ppExpected := PrintParseResult(expected)
		ppResult := PrintParseResult(result)
		t.Errorf("Expected \n%s\n, got \n%s", ppExpected, ppResult)
	}

	// Test from the center
	result, err = ParseExpression(input, 3)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expected = ParseResult{
		NextPosition: 8,
		Node: Node{
			Typ:   "CallExpression",
			Value: "subtract",
			Params: []Node{
				{Typ: "NumberLiteral", Value: "4", Params: []Node{}},
				{Typ: "NumberLiteral", Value: "2", Params: []Node{}},
			},
		},
	}
	if !reflect.DeepEqual(result, expected) {
		ppExpected := PrintParseResult(expected)
		ppResult := PrintParseResult(result)
		t.Errorf("Expected \n%s\n, got \n%s", ppExpected, ppResult)
	}
}

func TestParseProgram(t *testing.T) {

	input := []tokenize.Token{
		{Typ: "paren", Value: "("},
		{Typ: "symbol", Value: "print"},
		{Typ: "string", Value: "Hello"},
		{Typ: "number", Value: "2"},
		{Typ: "paren", Value: ")"},
		{Typ: "paren", Value: "("},
		{Typ: "symbol", Value: "add"},
		{Typ: "number", Value: "2"},
		{Typ: "paren", Value: "("},
		{Typ: "symbol", Value: "subtract"},
		{Typ: "number", Value: "4"},
		{Typ: "number", Value: "2"},
		{Typ: "paren", Value: ")"},
		{Typ: "paren", Value: ")"},
	}
	result, err := ParseProgram(input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expected := Node{
		Typ:   "Program",
		Value: "",
		Params: []Node{
			{
				Typ:   "CallExpression",
				Value: "print",
				Params: []Node{
					{Typ: "StringLiteral", Value: "Hello", Params: []Node{}},
					{Typ: "NumberLiteral", Value: "2", Params: []Node{}},
				},
			},
			{
				Typ:   "CallExpression",
				Value: "add",
				Params: []Node{
					{Typ: "NumberLiteral", Value: "2", Params: []Node{}},
					{
						Typ:   "CallExpression",
						Value: "subtract",
						Params: []Node{
							{Typ: "NumberLiteral", Value: "4", Params: []Node{}},
							{Typ: "NumberLiteral", Value: "2", Params: []Node{}},
						},
					},
				},
			},
		},
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %+v, got %+v", expected, result)
	}
}
