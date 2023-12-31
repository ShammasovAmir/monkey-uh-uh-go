package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	/* Identifiers + literals */
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 1343456
	/* Operators */
	ASSIGN = "="
	PLUS   = "+"
	/* Delimiters */
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	/* Keywords */
	FUNCTION = "FUNCTION"
	VAR      = "VAR"
)

var keywords = map[string]TokenType{
	"def": FUNCTION,
	"var": VAR,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
