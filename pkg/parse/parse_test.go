package parse

import (
	"reflect"
	"scheme2go/pkg/tokenize"
	"testing"
)

func TestParseNumber(t *testing.T) {
	input := []tokenize.Token{tokenize.Token{Typ: "Number", Value: "12345"}}
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
	input := Node{"StringLiteral", "abc", []Node{}}
	result, err := ParseString(input, 0)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expected := ParseResult{
		NextPosition: 1,
		Node: Node{
			Typ:    "StringLiteral",
			Value:  "abc",
			Params: []Node{},
		},
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %+v, got %+v", expected, result)
	}
}

func TestParseExpression(t *testing.T) {

}

func TestParseToken(t *testing.T) {
	input := "12345bhjkhuil"
	result, err := tokenizePattern("number", "[0-9]+", input, 0)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expected := MToken{5, Token{"number", "12345"}}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %+v, got %+v", expected, result)
	}
}
