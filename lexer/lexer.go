package lexer

import (
	"regexp"
	"strings"

	"github.com/six-nine/uzh/token"
)

type Lexer struct {
	bufferSize int64
}

func New() *Lexer {
	return new(Lexer)
}

func buildToken(tokType token.TokenType, literal string) token.Token {
	return token.Token{Type: tokType, Literal: literal}
}

func (lexer *Lexer) Tokenize(input string) []token.Token {
	input += " "

	var tokens []token.Token

	var buffer string

	resetBuffer := func() {
		if len(buffer) == 0 {
			return
		}

		tokType := token.IDENTIFIER

		tokenType, ok := token.Keywords[buffer]
		if ok {
			tokType = tokenType
		} else {
			var isIntLiteral, _ = regexp.MatchString(`\d+`, buffer)
			if isIntLiteral {
				tokType = token.INT_LITERAL
			}
			var isFloatLiteral, _ = regexp.MatchString(`\d+\.\d+`, buffer)
			if isFloatLiteral {
				tokType = token.FLOAT_LITERAL
			}
			var isCharLiteral, _ = regexp.MatchString(`\'.\'`, buffer)
			if isCharLiteral {
				tokType = token.CHAR_LITERAL
			}
			var isStringLiteral, _ = regexp.MatchString(`\"*\"`, buffer)
			if isStringLiteral {
				tokType = token.STRING_LITERAL
			}
		}

		tokens = append(
			tokens,
			buildToken(tokType, buffer),
		)

		buffer = ""
	}

	const Whitespace = " \n\t"

	for i := 0; i < len(input); i++ {
		if i+1 < len(input) {
			twoSymbolToken, ok := token.TwoSymbolTokens[input[i:i+2]]
			if ok {
				resetBuffer()
				tokens = append(tokens, buildToken(twoSymbolToken, input[i:i+2]))
				i++
				continue
			}
		}

		oneSymbolToken, ok := token.OneSymbolTokens[input[i:i+1]]
		if ok {
			resetBuffer()
			tokens = append(tokens, buildToken(oneSymbolToken, input[i:i+1]))
			continue
		}

		if strings.Contains(Whitespace, input[i:i+1]) {
			resetBuffer()
			continue
		}

		buffer += input[i : i+1]
	}

	return tokens
}
