package tkn

import (
	"fmt"
	"regexp"
)

type TokenType int

const (
	TokenTypeNil TokenType = iota
	TokenTypeNumber
	TokenTypeSymbol
	TokenTypeString
	TokenTypeParen
)

func (t TokenType) String() string {
	return [...]string{"Nil", "Number", "Symbol", "String", "Paren"}[t]
}

type TokenizeResult struct {
	ConsumedChars int
	Token         Token
}

type Token struct {
	Typ   TokenType
	Value string
}

func (t Token) IsNil() bool {
	return t.Typ == TokenTypeNil
}

func (t Token) String() string {
	// Assuming a maximum length of 10 for `typ`. Adjust the width as needed.
	return fmt.Sprintf("type: %-10s, value: %s", t.Typ, t.Value)
}

func TokenizeWhitespace(input string, current int) (TokenizeResult, error) {
	char := string(input[current])

	re, err := regexp.Compile(`\s`)
	if err != nil {
		return TokenizeResult{0, Token{}}, fmt.Errorf("failed to compile regex pattern: %s", err)
	}

	if re.MatchString(char) {
		return TokenizeResult{ConsumedChars: 1}, nil
	}

	return TokenizeResult{ConsumedChars: 0}, nil
}

func TokenizeCharacter(typ TokenType, value string, input string, current int) (TokenizeResult, error) {
	// check if beyond the end of the input
	if current >= len(input) {
		return TokenizeResult{0, Token{}}, nil
	}

	// check if the character matches
	if string(input[current]) == value {
		return TokenizeResult{1, Token{typ, value}}, nil
	}

	// no match
	return TokenizeResult{0, Token{}}, nil
}

func TokenizeParenOpen(input string, current int) (TokenizeResult, error) {
	return TokenizeCharacter(TokenTypeParen, "(", input, current)
}

func TokenizeParenClose(input string, current int) (TokenizeResult, error) {
	return TokenizeCharacter(TokenTypeParen, ")", input, current)
}

func TokenizePattern(typ TokenType, pattern string, input string, current int) (TokenizeResult, error) {
	// Check if beyond the end of the input
	if current >= len(input) {
		return TokenizeResult{0, Token{}}, nil
	}

	// Compile the regex pattern once outside the loop
	re, err := regexp.Compile(pattern)
	if err != nil {
		return TokenizeResult{0, Token{}}, fmt.Errorf("failed to compile regex pattern: %s", err)
	}

	// Keep eating characters until we find a non-match
	consumedChars := 0
	value := ""
	for current+consumedChars < len(input) && re.MatchString(string(input[current+consumedChars])) {
		value += string(input[current+consumedChars])
		consumedChars++
	}

	// Return the token if we consumed at least one character
	if consumedChars > 0 {
		return TokenizeResult{consumedChars, Token{typ, value}}, nil
	}

	return TokenizeResult{0, Token{}}, nil
}

func TokenizeNumber(input string, current int) (TokenizeResult, error) {
	return TokenizePattern(TokenTypeNumber, "[0-9]", input, current)
}

func TokenizeSymbol(input string, current int) (TokenizeResult, error) {
	return TokenizePattern(TokenTypeSymbol, "[a-z]", input, current)
}

func TokenizeString(input string, current int) (TokenizeResult, error) {
	// Check if beyond the end of the input
	if current >= len(input) {
		return TokenizeResult{0, Token{}}, nil
	}

	// Fail if the first character is not a quote
	if input[current] != '"' {
		return TokenizeResult{0, Token{}}, nil
	}

	// Process characters until the closing quote or end of string
	value := ""
	consumedChars := 1 // Start after the opening quote
	for current+consumedChars < len(input) && input[current+consumedChars] != '"' {
		value += string(input[current+consumedChars])
		consumedChars++
	}

	// Check if the loop ended without finding a closing quote
	if current+consumedChars >= len(input) {
		return TokenizeResult{0, Token{}}, fmt.Errorf("unterminated string at position %d", current)
	}

	// Include the closing quote in consumed characters
	return TokenizeResult{consumedChars + 1, Token{TokenTypeString, value}}, nil
}

func Tokenize(input string) ([]Token, error) {
	tokenizers := []func(string, int) (TokenizeResult, error){
		TokenizeWhitespace,
		TokenizeParenOpen,
		TokenizeParenClose,
		TokenizeNumber,
		TokenizeSymbol,
		TokenizeString,
	}

	current := 0
	tokens := []Token{}
	for current < len(input) {
		found := false
		for _, tokenizer := range tokenizers {
			mtoken, err := tokenizer(input, current)
			if err != nil {
				return nil, err
			}
			if mtoken.ConsumedChars > 0 {
				current += mtoken.ConsumedChars

				// check if token is null
				if mtoken.Token.IsNil() {
					continue
				}

				found = true
				tokens = append(tokens, mtoken.Token)
				break
			}
		}
		if !found {
			return nil, fmt.Errorf("no tokenizer found for input: %s", input[current:])
		}
	}

	return tokens, nil
}
