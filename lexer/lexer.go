package lexer

import (
	"errors"
	"regexp"
	"unicode"

	"github.com/six-nine/uzh/token"
)

type Lexer struct {
	buffer          string
	twoSymbolBuffer string
	tokens          []token.Token
}

func New() *Lexer {
	l := Lexer{"", "", []token.Token{}}
	return &l
}

func buildToken(tokType token.TokenType, literal string) token.Token {
	return token.Token{Type: tokType, Literal: literal}
}

func (lexer *Lexer) resetBuffer() {
	if len(lexer.buffer) == 0 {
		return
	}

	tokType := token.IDENTIFIER

	tokenType, ok := token.Keywords[lexer.buffer]
	if ok {
		tokType = tokenType
	} else {
		var isIntLiteral, _ = regexp.MatchString(`\b\d+\b`, lexer.buffer)
		if isIntLiteral {
			tokType = token.INT_LITERAL
		}
		var isFloatLiteral, _ = regexp.MatchString(`\b\d+\.\d+\b`, lexer.buffer)
		if isFloatLiteral {
			tokType = token.FLOAT_LITERAL
		}
		var isCharLiteral, _ = regexp.MatchString(`\b\'.\'\b`, lexer.buffer)
		if isCharLiteral {
			tokType = token.CHAR_LITERAL
		}
		var isStringLiteral, _ = regexp.MatchString(`\b\"*\"\b`, lexer.buffer)
		if isStringLiteral {
			tokType = token.STRING_LITERAL
		}
	}

	lexer.tokens = append(
		lexer.tokens,
		buildToken(tokType, lexer.buffer),
	)

	lexer.buffer = ""
}

func (lexer *Lexer) tryOneSymbolToken() error {
	oneSymbolToken, ok := token.OneSymbolTokens[lexer.twoSymbolBuffer[0:1]]
	if ok {
		lexer.resetBuffer()
		lexer.tokens = append(lexer.tokens, buildToken(oneSymbolToken, string(lexer.twoSymbolBuffer[0])))
		if unicode.IsSpace(rune(lexer.twoSymbolBuffer[1])) {
			lexer.twoSymbolBuffer = ""
		} else {
			lexer.twoSymbolBuffer = string(lexer.twoSymbolBuffer[1])
		}
		return nil
	} else {
		return errors.New("Not a one symbol token")
	}
}

func (lexer *Lexer) tryTwoSymbolToken() error {
	twoSymbolToken, ok := token.TwoSymbolTokens[lexer.twoSymbolBuffer]
	if ok {
		lexer.resetBuffer()
		lexer.tokens = append(lexer.tokens, buildToken(twoSymbolToken, lexer.twoSymbolBuffer))
		lexer.twoSymbolBuffer = ""
		return nil
	} else {
		return errors.New("Not a two symbol token")
	}
}

func (lexer *Lexer) AddChar(char byte) {
	cs := string(char)

	lexer.twoSymbolBuffer += cs

	if len(lexer.twoSymbolBuffer) == 2 {
		err := lexer.tryTwoSymbolToken()
		if err == nil {
			return
		}

		err = lexer.tryOneSymbolToken()
		if err == nil {
			return
		}

		lexer.buffer += string(lexer.twoSymbolBuffer[0])
		lexer.twoSymbolBuffer = string(lexer.twoSymbolBuffer[1])
	}

	if unicode.IsSpace(rune(char)) {
		lexer.twoSymbolBuffer = ""
		lexer.resetBuffer()
	}
}

func (lexer *Lexer) ExtractTokens() []token.Token {
	if len(lexer.twoSymbolBuffer) == 1 {
		lexer.twoSymbolBuffer += " "
		err := lexer.tryOneSymbolToken()
		if err != nil {
			lexer.buffer += lexer.twoSymbolBuffer
			lexer.twoSymbolBuffer = ""
			lexer.resetBuffer()
		}
	}
	tokens := lexer.tokens
	lexer.tokens = lexer.tokens[:0]
	return tokens
}
