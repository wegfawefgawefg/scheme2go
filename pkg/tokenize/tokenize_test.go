package tokenize

import (
	"reflect"
	"testing"
)

func TestTokenizeCharacter(t *testing.T) {
	input := "example"
	result, err := TokenizeCharacter("char", "e", input, 0)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expected := MToken{1, Token{"char", "e"}}
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
	expected := MToken{1, Token{"paren", "("}}
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
	expected := MToken{1, Token{"paren", ")"}}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %+v, got %+v", expected, result)
	}
}

func TestTokenizePattern(t *testing.T) {
	input := "12345bhjkhuil"
	result, err := TokenizePattern("number", "[0-9]+", input, 0)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expected := MToken{5, Token{"number", "12345"}}
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
	expected := MToken{3, Token{"number", "123"}}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %+v, got %+v", expected, result)
	}
}

func TestTokenizeName(t *testing.T) {
	input := "hello world"
	result, err := TokenizeName(input, 0)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expected := MToken{5, Token{"name", "hello"}}
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
	expected := MToken{13, Token{"string", "hello world"}}
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
	expected := MToken{0, Token{}}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %+v for unterminated string, got %+v", expected, result)
	}
}

func TestTokenizePatternInvalidRegex(t *testing.T) {
	input := "12345"
	result, err := TokenizePattern("number", "[0-9", input, 0) // Invalid regex pattern
	if err == nil {
		t.Errorf("Expected error for invalid regex, but got nil")
	}
	expected := MToken{0, Token{}}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %+v for invalid regex, got %+v", expected, result)
	}
}
