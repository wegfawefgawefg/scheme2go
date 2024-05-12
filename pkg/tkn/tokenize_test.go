package tkn

import (
	"reflect"
	"testing"
)

/*
The token type is nil, bc its not really a token. It's just a character.
*/
func TestTokenizeCharacter(t *testing.T) {
	input := "example"
	result, err := TokenizeCharacter(TokenTypeNil, "e", input, 0)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expected := TokenizeResult{1, Token{TokenTypeNil, "e"}}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %+v, got %+v", expected, result)
	}
}

func TestTokenizeParenOpen(t *testing.T) {
	input := "("
	result, err := TokenizeParenOpen(input, 0)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expected := TokenizeResult{1, Token{TokenTypeParen, "("}}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %+v, got %+v", expected, result)
	}
}

func TestTokenizeParenClose(t *testing.T) {
	input := ")"
	result, err := TokenizeParenClose(input, 0)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expected := TokenizeResult{1, Token{TokenTypeParen, ")"}}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %+v, got %+v", expected, result)
	}
}

func TestTokenizePattern(t *testing.T) {
	input := "12345bhjkhuil"
	result, err := TokenizePattern(TokenTypeNumber, "[0-9]+", input, 0)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expected := TokenizeResult{5, Token{TokenTypeNumber, "12345"}}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %+v, got %+v", expected, result)
	}
}

func TestTokenizeNumber(t *testing.T) {
	input := "123abc"
	result, err := TokenizeNumber(input, 0)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expected := TokenizeResult{3, Token{TokenTypeNumber, "123"}}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %+v, got %+v", expected, result)
	}
}

func TestTokenizeSymbol(t *testing.T) {
	input := "hello world"
	result, err := TokenizeSymbol(input, 0)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expected := TokenizeResult{5, Token{TokenTypeSymbol, "hello"}}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %+v, got %+v", expected, result)
	}
}

func TestTokenizeString(t *testing.T) {
	input := `"hello world"`
	result, err := TokenizeString(input, 0)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expected := TokenizeResult{13, Token{TokenTypeString, "hello world"}}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %+v, got %+v", expected, result)
	}
}

func TestTokenizeStringWithNoClosingQuote(t *testing.T) {
	input := `"hello world`
	result, err := TokenizeString(input, 0)
	if err == nil {
		t.Errorf("Expected error for unterminated string, but got nil")
	}
	expected := TokenizeResult{0, Token{}}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %+v for unterminated string, got %+v", expected, result)
	}
}

func TestTokenizePatternInvalidRegex(t *testing.T) {
	input := "12345"
	result, err := TokenizePattern(TokenTypeNumber, "[0-9", input, 0) // Invalid regex pattern
	if err == nil {
		t.Errorf("Expected error for invalid regex, but got nil")
	}
	expected := TokenizeResult{0, Token{}}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %+v for invalid regex, got %+v", expected, result)
	}
}
