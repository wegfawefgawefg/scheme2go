package tokenize

import (
	"fmt"
	"regexp"
)

type MToken struct {
	ConsumedChars int
	Token         Token
}

type Token struct {
	Typ   string
	Value string
}

func (t Token) IsNil() bool {
	return t.Typ == "" && t.Value == ""
}

func (t Token) String() string {
	// Assuming a maximum length of 10 for `typ`. Adjust the width as needed.
	return fmt.Sprintf("type: %-10s, value: %s", t.Typ, t.Value)
}

func ItsWhitespace(input string, current int) (MToken, error) {
	char := string(input[current])

	re, err := regexp.Compile(`\s`)
	if err != nil {
		return MToken{0, Token{}}, fmt.Errorf("failed to compile regex pattern: %s", err)
	}

	if re.MatchString(char) {
		return MToken{ConsumedChars: 1}, nil
	}

	return MToken{ConsumedChars: 0}, nil
}

func TokenizeCharacter(typ string, value string, input string, current int) (MToken, error) {
	// check if beyond the end of the input
	if current >= len(input) {
		return MToken{0, Token{}}, nil
	}

	// check if the character matches
	if string(input[current]) == value {
		return MToken{1, Token{typ, value}}, nil
	}

	// no match
	return MToken{0, Token{}}, nil
}

func TokenizeParenOpen(input string, current int) (MToken, error) {
	return TokenizeCharacter("paren", "(", input, current)
}

func TokenizeParenClose(input string, current int) (MToken, error) {
	return TokenizeCharacter("paren", ")", input, current)
}

func TokenizePattern(typ string, pattern string, input string, current int) (MToken, error) {
	// Check if beyond the end of the input
	if current >= len(input) {
		return MToken{0, Token{}}, nil
	}

	// Compile the regex pattern once outside the loop
	re, err := regexp.Compile(pattern)
	if err != nil {
		return MToken{0, Token{}}, fmt.Errorf("failed to compile regex pattern: %s", err)
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
		return MToken{consumedChars, Token{typ, value}}, nil
	}

	return MToken{0, Token{}}, nil
}

func TokenizeNumber(input string, current int) (MToken, error) {
	return TokenizePattern("number", "[0-9]", input, current)
}

func TokenizeName(input string, current int) (MToken, error) {
	return TokenizePattern("name", "[a-z]", input, current)
}

func TokenizeString(input string, current int) (MToken, error) {
	// Check if beyond the end of the input
	if current >= len(input) {
		return MToken{0, Token{}}, nil
	}

	// Fail if the first character is not a quote
	if input[current] != '"' {
		return MToken{0, Token{}}, nil
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
		return MToken{0, Token{}}, fmt.Errorf("unterminated string at position %d", current)
	}

	// Include the closing quote in consumed characters
	return MToken{consumedChars + 1, Token{"string", value}}, nil
}

func Tokenize(input string) ([]Token, error) {
	tokenizers := []func(string, int) (MToken, error){
		ItsWhitespace,
		TokenizeParenOpen,
		TokenizeParenClose,
		TokenizeNumber,
		TokenizeName,
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
