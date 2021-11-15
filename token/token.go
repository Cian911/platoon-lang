package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Iddentifiers & Literals
	IDENT = "IDENT"
	INT   = "INT"

	// Operators
	ASSIGN  = "="
	PLUS    = "+"
	MINUS   = "-"
	BANG    = "!"
	ASTERIX = "*"
	EQ      = "=="
	NOT_EQ  = "!="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	SLASH     = "/"

	LPARAN = "("
	RPARAN = ")"
	LBRACE = "{"
	RBRACE = "}"
	GT     = ">"
	LT     = "<"

	// Keyswords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

var languageKeywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := languageKeywords[ident]; ok {
		return tok
	}

	return IDENT
}
