package token

// TokenType is the base of all tokens to define their actual type
type TokenType string

// Token is the prime structure of a token
type Token struct {
	Type TokenType
	Literal string
}

const (
	// ILLEGAL is when a token is not correct
	ILLEGAL = "ILLEGAL"
	// EOF end of the tokenized source file
	EOF = "EOF"

	// IDENT is an identifier token
	IDENT = "IDENT"

	// ASSIGN is an assignment token
	ASSIGN = "ASSIGN"

	// COMMA is a comma sign token
	COMMA = "COMMA"

	// SEMICOLON to finish an expression or a statement
	SEMICOLON = "SEMICOLON"

	SEND = "->"
	RETURN = "=>"

	// LPAREN is a left parenthese token
	LPAREN = "("
	// RPAREN is a right parenthese token
	RPAREN = ")"
	// LBRACE is a left brace token
	LBRACE = "{"
	// RBRACE is a right brace token
	RBRACE = "}"

	// LET is a start of a definition
	LET = "LET"

	// DIAGRAM is the enclosing declaration of a real diagram
	DIAGRAM = "DIAGRAM"

	INCLUDE = "INCLUDE"

	STRING = "STRING"
)

var keywords = map[string]TokenType {
	"diagram": DIAGRAM,
	"include": INCLUDE,
}

// LookupIdent checks the keyword table to see whether the given identifier is in fact a keyword
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}

