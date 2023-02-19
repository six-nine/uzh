package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	IDENTIFIER TokenType = "IDENTIFIER" // x, y, aboba, _x1

	// literals
	INT_LITERAL    TokenType = "INT_LITERAL"
	FLOAT_LITERAL  TokenType = "FLOAT_LITERAL"
	CHAR_LITERAL   TokenType = "CHAR_LITERAL"
	STRING_LITERAL TokenType = "STRING_LITERAL"
	TRUE           TokenType = "TRUE"
	FALSE          TokenType = "FALSE"

	// data types
	INT    TokenType = "INT_TYPE"
	FLOAT  TokenType = "FLOAT_TYPE"
	CHAR   TokenType = "CHAR_TYPE"
	STRING TokenType = "STRING_TYPE"
	BOOL   TokenType = "BOOL_TYPE"

	// keywords
	PROCEDURE TokenType = "PROCEDURE"
	MAIN      TokenType = "MAIN"
	IF        TokenType = "IF"
	ELSE      TokenType = "ELSE"
	RETURN    TokenType = "RETURN"

	ASSIGN    TokenType = "="
	PLUS      TokenType = "+"
	MINUS     TokenType = "-"
	ASTERISK  TokenType = "*"
	SLASH     TokenType = "/"
	COMMA     TokenType = ","
	SEMICOLON TokenType = ";"

	EQUAL     TokenType = "=="
	LT        TokenType = "<"
	GT        TokenType = ">"
	L_OR_EQ   TokenType = "<="
	G_OR_EG   TokenType = ">="
	NOT_EQUAL TokenType = "!="
)
