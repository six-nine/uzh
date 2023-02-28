package lexer

import (
	"errors"

	"github.com/six-nine/uzh/token"
)

var ErrEOF = errors.New("End of file reached")

type CodeSource interface {
	GetChar() (byte, error) // should return char or error
}

type Lexer struct {
	currentChar byte
	codeSource  *CodeSource
}

func New(codeSource *CodeSource) *Lexer {
	l := Lexer{}
	l.codeSource = codeSource
	return &l
}

func buildToken(tokType token.TokenType, literal string) token.Token {
	return token.Token{Type: tokType, Literal: literal}
}

func isNumber(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isLetterOrUnderscore(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || ch == '_'
}

func (lexer *Lexer) getIdentifierOrKeyword() token.Token {

}

func (lexer *Lexer) getIntOrFloatLiteral() token.Token {

}

func (lexer *Lexer) getStringLiteral() token.Token {

}

func (lexer *Lexer) getCharLiteral() token.Token {
	literal := string(lexer.currentChar)
	ch, err := lexer.codeSource.GetChar()
	literal += string(ch)
	if ch == '\\' {
		// escape char
	} else {

	}
}

func (lexer *Lexer) getOneOrTwoSymbolOperator() token.Token {
	first_ch := lexer.currentChar
	lexer.currentChar = (*lexer.codeSource).GetChar()

	twoSymbOp := string(first_ch) + string(lexer.currentChar)
	tokType, isTwoSymbToken := token.TwoSymbolTokens[twoSymbOp]
	if isTwoSymbToken {
		lexer.currentChar = 0
		return token.Token{Type: tokType, Literal: twoSymbOp}
	}

	oneSymbOp := string(first_ch)
	tokType, isOneSymbToken := token.OneSymbolTokens[oneSymbOp]
	if isOneSymbToken {
		return token.Token{Type: tokType, Literal: oneSymbOp}
	}

	return token.Token{Type: token.ILLEGAL, Literal: ""}
}

func (lexer *Lexer) GetToken() (token.Token, error) {
	// if no char from previous step, read it
	if lexer.currentChar == 0 {
		ch, err := (*lexer.codeSource).GetChar()
		lexer.currentChar = ch
		if err == ErrEOF {
			return token.Token{token.END_OF_PROGRAM, ""}, nil
		} else if err != nil {
			return err
		}
	}

	// if not read, then end of file/line
	if lexer.currentChar == 0 {
	}

	if isLetterOrUnderscore(lexer.currentChar) {
		return lexer.getIdentifierOrKeyword()
	}

	if isNumber(lexer.currentChar) {
		return lexer.getIntOrFloatLiteral()
	}

	if lexer.currentChar == '"' {
		return lexer.getStringLiteral()
	}

	if lexer.currentChar == '\'' {
		return lexer.getCharLiteral()
	}

	return lexer.getOneOrTwoSymbolOperator() // try to get operator, return illegal otherwise
}
