package token

// TokenType Let's have a formal type for tokens. In real-world, they can be much more efficient,
// but let's have it as string for now. We can also easily print it which is useful for debugging.
type TokenType string

type Token struct {
	Type    TokenType
	Literal string // this is the literal value of the token
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 1343456
	// Operators
	ASSIGN = "="
	PLUS   = "+"
	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
