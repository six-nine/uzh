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

	LT        TokenType = "<"
	GT        TokenType = ">"
	L_OR_EQ   TokenType = "<="
	G_OR_EG   TokenType = ">="
	NOT_EQUAL TokenType = "!="
	EQUAL     TokenType = "=="

	LPAREN   TokenType = "("
	RPAREN   TokenType = ")"
	LBRACE   TokenType = "{"
	RBRACE   TokenType = "}"
	LBRACKET TokenType = "["
	RBRACKET TokenType = "]"
)

var Keywords = map[string]TokenType{
	"int":    INT,
	"float":  FLOAT,
	"char":   CHAR,
	"string": STRING,
	"bool":   BOOL,
	"proc":   PROCEDURE,
	"main":   MAIN,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

var OneSymbolTokens = map[string]TokenType{
	"=": ASSIGN,
	"+": PLUS,
	"-": MINUS,
	"*": ASTERISK,
	"/": SLASH,
	",": COMMA,
	";": SEMICOLON,
	"<": LT,
	">": GT,
	"(": LPAREN,
	")": RPAREN,
	"{": LBRACE,
	"}": RBRACE,
	"[": LBRACKET,
	"]": RBRACKET,
}

var TwoSymbolTokens = map[string]TokenType{
	"<=": L_OR_EQ,
	">=": G_OR_EG,
	"==": EQUAL,
	"!=": NOT_EQUAL,
}
